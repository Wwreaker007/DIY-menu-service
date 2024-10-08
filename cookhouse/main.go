package main

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClientConnection(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error in creating a client connection to order service")
		return nil, err
	}
	return conn, nil
}

func main() {
	/*
		1. Open a connection to the port for starting the HTTP server
		2. Start the HTTP server registering the service implementation
	*/
	// ordersUrl := os.Getenv("ORDERS_HOST") + ":" + os.Getenv("ORDERS_PORT")
	ordersUrl := "orders:9001"
	conn, err := NewGRPCClientConnection(ordersUrl)
	if err != nil {
		fmt.Println("error : ", err.Error())
		panic(0)
	}
	defer conn.Close()
	httpServer := NewCookHouseHttpServer("tcp", ":8000", conn)
	err = httpServer.Start()
	if err != nil {
		fmt.Println("error : ", err.Error())
		panic(0)
	}
}