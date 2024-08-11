package types

import (
	"context"

	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type CookHouse interface {
	GetAllOrdersByStatusFilter(context.Context, *data.GetOrderByStatusFilterRequest) (*data.GetOrderByStatusFilterResponse, error)
	GetOrderByOrderID(context.Context, *data.GetOrderByOrderIDRequest) (*data.GetOrderByOrderIDResponse, error)
	UpdateOrderStatus(context.Context, *data.UpdateOrderStatusRequest) (*data.UpdateOrderStatusResponse, error)
}