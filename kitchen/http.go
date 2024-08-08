package main

import (
	"fmt"
	"net"
)

type HttpServer struct {
	Network		string
	Address 	string
}

func NewHttpServer(network string, address string) *HttpServer{
	return &HttpServer{
		Network: network,
		Address: address,
	}
}

func (s *HttpServer) Start() (net.Listener, error) {
	connection, err := net.Listen(s.Network, s.Address)
	if err != nil {
		fmt.Println("error listening on port 9000")
		return nil, err
	}
	return connection, nil
}