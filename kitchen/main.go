package main

import (
	"fmt"
)

func main() {
	/*
		1. Open a connection to the port for starting the HTTP server
		2. Start the HTTP server registering the service implementation
	*/

	httpServer := NewHttpServer("tcp", ":9000")
	_, err := httpServer.Start()
	if err != nil {
		fmt.Println("Port 9000 is not available")
		panic(0)
	}
}