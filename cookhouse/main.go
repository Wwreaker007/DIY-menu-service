package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClientConnection(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error in creating a client connection to order service")
	}
	return conn
}

func main() {
	/*
		1. Open a connection to the port for starting the HTTP server
		2. Start the HTTP server registering the service implementation
	*/
	conn := NewGRPCClientConnection(":9001")
	defer conn.Close()
	httpServer := NewCookHouseHttpServer("tcp", ":8000", conn)
	err := httpServer.Start()
	if err != nil {
		fmt.Println("error : ", err.Error())
		panic(0)
	}
}