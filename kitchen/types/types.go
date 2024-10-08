package types

import (
	"context"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type Kitchen interface {
	CreateOrder(context.Context, *data.CreateOrderRequest) (*data.CreateOrderResponse, error)
	GetOrder(context.Context, *data.GetOrderRequest) (*data.GetOrderResponse, error)
	UpdateOrder(context.Context, *data.UpdateOrderRequest) (*data.UpdateOrderResponse, error)
}