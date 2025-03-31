package Utility

import (
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
}

func TerminateSocketConnection() {
	SocketConnection.Close()
}

func Sender() {}

func Listener() {
	for {
		var buffer []byte
		_, err := SocketConnection.Read(buffer)
		if err != nil {
			log.Print(err)
		}
		myString := string(buffer[:])

		if myString == "HeartBeat" {
			Sender()
		}

	}
}
