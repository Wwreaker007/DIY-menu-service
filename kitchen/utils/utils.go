package utils

import (
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

func TransformToGRPCCreateOrderRequest(source *data.CreateOrderRequest) (*orders.CreateOrderRequest) {
	return &orders.CreateOrderRequest{
		UserID: source.UserID,
		Order: source.Order,
	}
}

func TransformToHTTPCreateOrderResponse(source *orders.CreateOrderResponse) (*data.CreateOrderResponse) {
	return &data.CreateOrderResponse{
		Status: source.GetStatus(),
		Order: source.GetOrder(),
	}
}

func TransformToGRPCGetOrderRequest(source *data.GetOrderRequest) (*orders.GetOrderRequest) {
	return &orders.GetOrderRequest{
		UserID: source.UserID,
		OrderStatus: source.OrderStatus,
	}
}

func TransformToHTTPGetOrderRespone(source *orders.GetOrderResponse) (*data.GetOrderResponse) {
	return &data.GetOrderResponse{
		Status: source.GetStatus(),
		Orders: source.GetOrders(),
	}
}

func TransformToGRPCUpdateOrderRequest(source *data.UpdateOrderRequest) (*orders.UpdateOrderRequest) {
	return &orders.UpdateOrderRequest{
		UserID: source.UserID,
		Order: source.Order,
	}
}

func TransformToHTTPUpdateOrderResponse(source *orders.UpdateOrderResponse) (*data.UpdateOrderResponse) {
	return &data.UpdateOrderResponse{
		Status: source.GetStatus(),
		Order: source.GetOrder(),
	}
}