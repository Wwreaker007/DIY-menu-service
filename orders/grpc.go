package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/Wwreaker007/DIY-menu-service/orders/db/postgres"
	"github.com/Wwreaker007/DIY-menu-service/orders/handlers"
	cms "github.com/Wwreaker007/DIY-menu-service/orders/services/cookhouse_manager"
	oms "github.com/Wwreaker007/DIY-menu-service/orders/services/order_manager"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Network 	string
	Address 	string
	dbclient	*sql.DB
}

func NewGrpcServer(network string, address string, client *sql.DB) *GrpcServer {
	return &GrpcServer{
		Network: network,
		Address: address,
		dbclient: client,
	}
}

func (s *GrpcServer) Start() error {
	connection, err := net.Listen(s.Network, s.Address)
	if err != nil {
		fmt.Println("error :", err.Error())
		return err
	}

	/*
		Spin up the dependencies required for the GRPC server
		1. Replacing the inmemoryDB implementation with postgresDB CRUD. 
	*/

	// db := inmem.NewInMemoryDbService(data.NewThreadSafeOrderEntity())
	db := postgres.NewPostgresDBService(s.dbclient)
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