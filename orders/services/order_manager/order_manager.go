package ordermanager

import (
	"context"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type OrderManagerService struct {
	orders.UnimplementedOrderServiceServer	
}

func NewOrderManagerService() *OrderManagerService {
	return &OrderManagerService{}
}

func (oms *OrderManagerService) CreateOrder(ctx context.Context, request *orders.CreateOrderRequest) (response *orders.CreateOrderResponse, err error) {

	return response, err
}


func (oms *OrderManagerService) GetOrder(ctx context.Context, request *orders.GetOrderRequest) (response *orders.GetOrderResponse, err error) {
	
	return response, err
}

func (oms *OrderManagerService) UpdateOrder(ctx context.Context, request *orders.UpdateOrderRequest) (response *orders.UpdateOrderResponse, err error) {

	return response, err
}