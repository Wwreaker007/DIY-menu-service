package handlers

import (
	"context"
	"log"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/orders/types"
	"google.golang.org/grpc"
)

type OrderManagerHandler struct {
	OrderManagerService types.Order
	orders.UnimplementedOrderServiceServer
}

func NewOrderManagerhandler(server *grpc.Server, orderManagerService types.Order) {
	orderManagerHandler := &OrderManagerHandler{
		OrderManagerService: orderManagerService,
	}

	/*
		1. Get the service implementation to the handler layer
		2. Register this handler to the OrderGRPCServer
	*/
	orders.RegisterOrderServiceServer(server, orderManagerHandler)
}

func (h *OrderManagerHandler) CreateOrder(ctx context.Context, request *orders.CreateOrderRequest) (*orders.CreateOrderResponse,error) {
	response, err := h.OrderManagerService.CreateOrder(ctx, request)
	if err != nil {
		log.Println("Error in creating an order : ", err.Error())
		return nil, err
	}
	return response, err
}

func (h *OrderManagerHandler) GetOrder(ctx context.Context, request *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	response, err := h.OrderManagerService.GetOrder(ctx, request)
	if err != nil {
		log.Println("Error in creating an order : ", err.Error())
		return nil, err
	}
	return response, err
}

func (h *OrderManagerHandler) UpdateOrder(ctx context.Context, request *orders.UpdateOrderRequest) (*orders.UpdateOrderResponse, error) {
	response, err := h.OrderManagerService.UpdateOrder(ctx, request)
	if err != nil {
		log.Println("Error in creating an order : ", err.Error())
		return nil, err
	}
	return response, err
}