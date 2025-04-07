package Utility

import (
	"DistributedFileDBChunkServer/Helper/MetaData"
	GRPCHandler "DistributedFileDBChunkServer/Package/GRPC"
	"context"
	"log"
	"strings"

	"google.golang.org/grpc"
)

var SendMessage GRPCHandler.MessageExchangeServiceClient

func InitiateGRPCConnection() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Print(err)
	}
	messageService := GRPCHandler.NewMessageExchangeServiceClient(conn)

	SendMessage = messageService

	if len(MetaData.ChunkNameList) > 0 {
		Send()
	}

}

func Send() {
	message := GRPCHandler.RecievedMessage{}

	if MetaData.UsedSize > 0 {
		message.AvailableSize = string(MetaData.UsedSize)
	} else {
		message.AvailableSize = string(MetaData.TotalSize)
	}

	message.ServerID = MetaData.ServerID
	message.ChunkList = strings.Join(MetaData.ChunkNameList, ",")

	response, err := SendMessage.MessageProcessor(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err, response)
	}
}
