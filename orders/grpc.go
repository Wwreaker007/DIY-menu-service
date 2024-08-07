package orders

import (
	"fmt"
	"net"
)

type GrpcServer struct {
	Network string
	Address string
}

func NewGrpcServer(address string, network string) *GrpcServer {
	return &GrpcServer{
		Network: network,
		Address: address,
	}
}

func (s *GrpcServer) GetConnection() (net.Listener, error) {
	connection, err := net.Listen(s.Network, s.Address)
	if err != nil {
		fmt.Println("error listening on port 9001")
		return nil, err
	}

	return connection, nil
}