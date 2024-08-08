package main

import (
	"fmt"
)

func main () {
	/*
		1. Open a connection to the port for starting the GRPC server
		2. Start the GRPC server registering the service implementation
	*/
	grpcServer := NewGrpcServer("tcp", ":9001")
	err := grpcServer.Start()
	if err != nil {
		fmt.Println("Port 9001 is not available")
		panic(0)
	}
}