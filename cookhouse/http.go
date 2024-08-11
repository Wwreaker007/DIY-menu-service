package main

import (
	"log"
	"net/http"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/cookhouse"
	"github.com/Wwreaker007/DIY-menu-service/cookhouse/handlers"
	cookhousemanager "github.com/Wwreaker007/DIY-menu-service/cookhouse/services/cookhouse_manager"
	"google.golang.org/grpc"
)

const (
	GET_ALL_ORDER = "POST /order/all/v1"
	GET_ORDER_ORDERID = "POST /order/orderid/v1"
	UPDATE_ORDER_STATUS = "POST /update/order/status/v1"
)

type CookHouseHttpServer struct {
	Network 	string
	Address 	string
	Conn    	*grpc.ClientConn
}

func NewCookHouseHttpServer(network string, address string, conn *grpc.ClientConn) *CookHouseHttpServer {
	return &CookHouseHttpServer{
		Network: network,
		Address: address,
		Conn: conn,
	}
}

func (s *CookHouseHttpServer) Start() error {
	router := http.NewServeMux()

	// Create the service dependencies here to be used in handlers
	cms := cookhousemanager.NewCookhouseManagerService(cookhouse.NewCookHouseServiceClient(s.Conn))

	// Register the handler to the route
	router.HandleFunc(GET_ALL_ORDER, handlers.NewGetAllOrderByStatusFilterHandler(cms).Handler)
	router.HandleFunc(GET_ORDER_ORDERID, handlers.NewGetOrderByOrderIDHandler(cms).Handler)
	router.HandleFunc(UPDATE_ORDER_STATUS, handlers.NewUpdateOrderStatusHandler(cms).Handler)

	// Starting cookhouse HTTP service
	log.Println("Starting Cookhouse HTTP server on port: " + s.Address)
	return http.ListenAndServe(s.Address, router)
}