package main

import (
	"log"
	"net/http"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/kitchen/handlers"
	kitchenmanager "github.com/Wwreaker007/DIY-menu-service/kitchen/services/kitchen_manager"
	"google.golang.org/grpc"
)

const (
	CREATE_ORDER = "POST /create/order/v1"
	GET_ORDER = "POST /get/order/v1"
	UPDATE_ORDER = "POST /update/order/v1"
)

type HttpServer struct {
	Network		string
	Address 	string
	conn 		*grpc.ClientConn
}

func NewHttpServer(network string, address string, conn *grpc.ClientConn) *HttpServer{
	return &HttpServer{
		Network: network,
		Address: address,
		conn: conn,
	}
}

func (s *HttpServer) Start()  error {
	router := http.NewServeMux()

	// Create gRPC client for Order gRPC service
	client := orders.NewOrderServiceClient(s.conn)
	kms := kitchenmanager.NewKitchenManagerService(client)

	// Register the routes for the specific handlers
	router.HandleFunc(CREATE_ORDER, handlers.NewCreateOrderHandler(kms).CreatOrderHandler)
	router.HandleFunc(GET_ORDER, handlers.NewGetOrderHandler(kms).GetOrderHandler)
	router.HandleFunc(UPDATE_ORDER, handlers.NewUpdateOrderHandler(kms).UpdateOrderHandler)

	// Start the server on the passed on port
	log.Println("Starting HTTP server at port 9000")
	return http.ListenAndServe(s.Address, router)
}