package Utility

import (
	"DistributedFileDBChunkServer/Helper/MetaData"
	"DistributedFileDBChunkServer/Helper/StructStore"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

var SocketConnection net.Conn

func InitiateSocketConnection() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	SocketConnection = conn
	Sender()
	go recieve(SocketConnection)
}

func recieve(conn net.Conn) {
	rd := bufio.NewReader(conn)
	var buf []byte
	for {
		n, err := rd.Read(buf[:])
		if err != nil {
			fmt.Print("Reciever Error :")
			fmt.Print(err, n)
		}
		buf = bytes.Trim(buf, "\x00")
		recievedMessage := string(buf)
		if recievedMessage == "HeartBeat" {
			fmt.Printf("RecievedMessge: " + recievedMessage)
			Sender()
		}
	}
}

func Sender() {
	SendInstance := StructStore.ChunkMapping{}
	if MetaData.UsedSize > 0 {
		SendInstance.AvailableSpace = MetaData.UsedSize
	} else {
		SendInstance.AvailableSpace = MetaData.TotalSize
	}
	SendInstance.ServerID = MetaData.ServerID
	if MetaData.ChunkNameList == nil {
		MetaData.ChunkNameList = make([]string, 10)
	}
	SendInstance.ChunkList = MetaData.ChunkNameList
	SendMessage, err := json.Marshal(SendInstance)
	if err != nil {
		panic(err)
	}
	_, NewErr := SocketConnection.Write(SendMessage)
	if NewErr != nil {
		panic(NewErr)
	}
}

func Listener() {
	for {
		var buffer []byte
		_, err := SocketConnection.Read(buffer)
		if err != nil {
			log.Print(err)
		}
		if buffer != nil {
			myString := string(buffer[:])

			if myString == "HeartBeat" {
				Sender()
			}
		}

	}
}

func TerminateSocketConnection() {
	SocketConnection.Close()
}
