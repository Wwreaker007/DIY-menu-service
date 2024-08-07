package main

import (
	"fmt"

	"github.com/Wwreaker007/DIY-menu-service/kitchen"
	"github.com/Wwreaker007/DIY-menu-service/orders"
)

func main() {
	/*
		1. Open a connection to the port for starting the HTTP server
		2. Start the HTTP server registering the service implementation
	*/

	httpServer := kitchen.NewHttpServer("tcp", ":9000")
	_, err := httpServer.GetConnection()
	if err != nil {
		fmt.Println("Port 9000 is not available")
		panic(0)
	}

	

	/*
		1. Open a connection to the port for starting the GRPC server
		2. Start the GRPC server registering the service implementation
	*/
	grpcServer := orders.NewGrpcServer("tcp", ":9001")
	_, err = grpcServer.GetConnection()
	if err != nil {
		fmt.Println("Port 9001 is not available")
		panic(0)
	}


}