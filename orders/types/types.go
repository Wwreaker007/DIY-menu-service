package types

import (
	"context"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type Order interface {
	CreateOrder(context.Context, *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error)
	GetOrder(context.Context, *orders.GetOrderRequest) (*orders.GetOrderResponse, error)
	UpdateOrder(context.Context, *orders.UpdateOrderRequest) (*orders.UpdateOrderResponse, error)
}

type OrderRepositoryManager interface {
	UpsertOrder(context.Context, *data.OrderEntity) (error)
	GetAllOrdersByUserID(context.Context, string) ([]*data.OrderEntity, error)
	GetOrderByOrderID(context.Context, string) (*data.OrderEntity, error)
}