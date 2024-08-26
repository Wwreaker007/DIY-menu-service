package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Wwreaker007/DIY-menu-service/common/data"
	"github.com/Wwreaker007/DIY-menu-service/orders/db/inmem"
	"github.com/Wwreaker007/DIY-menu-service/orders/handlers"
	oms "github.com/Wwreaker007/DIY-menu-service/orders/services/order_manager"
	cms "github.com/Wwreaker007/DIY-menu-service/orders/services/cookhouse_manager"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Network string
	Address string
}

func NewGrpcServer(network string, address string) *GrpcServer {
	return &GrpcServer{
		Network: network,
		Address: address,
	}
}

func (s *GrpcServer) Start() error {
	connection, err := net.Listen(s.Network, s.Address)
	if err != nil {
		fmt.Println("error :", err.Error())
		return err
	}

	// Spin up the dependencies required for the GRPC server
	db := inmem.NewInMemoryDbService(data.NewThreadSafeOrderEntity())
	orderService := oms.NewOrderManagerService(db)
	cookhouseService := cms.NewCookHouseManagerService(db)

	// Create a new gRPC server which will serve the requests
	grpcServer := grpc.NewServer()

	// Register the services via the handlers here to the gRPC
	handlers.NewOrderManagerhandler(grpcServer, orderService)
	handlers.NewCookHouseManagerHandler(grpcServer, cookhouseService)

	log.Println("Starting order manager GRPC server ")
	return grpcServer.Serve(connection)
}