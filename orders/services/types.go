package services

import (
	"context"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
)

type Order interface {
	CreateOrder(context.Context, *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error)
	GetOrder(context.Context, *orders.GetOrderRequest) (*orders.GetOrderResponse, error)
	UpdateOrder(context.Context, *orders.UpdateOrderRequest) (*orders.UpdateOrderResponse, error)
}