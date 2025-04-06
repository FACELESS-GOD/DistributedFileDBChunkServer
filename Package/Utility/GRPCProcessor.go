package Utility

import (
	GRPCHandler "DistributedFileDBChunkServer/Package/GRPC"
	"log"

	"google.golang.org/grpc"
)

func InitiateGRPCConnection() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("",grpc.WithInsecure())
	if err !=  nil {
		log.Print(err)
	}
	messageService := GRPCHandler.NewMessageExchangeServiceClient(conn)

	
}