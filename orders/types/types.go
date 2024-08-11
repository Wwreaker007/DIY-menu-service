package types

import (
	"context"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/cookhouse"
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type Order interface {
	CreateOrder(context.Context, *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error)
	GetOrder(context.Context, *orders.GetOrderRequest) (*orders.GetOrderResponse, error)
	UpdateOrder(context.Context, *orders.UpdateOrderRequest) (*orders.UpdateOrderResponse, error)
}

type CookHouse interface {
	GetAllOrderByStatusFilter(context.Context, *cookhouse.GetAllOrderByStatusFilterRequest) (*cookhouse.GetAllOrderByStatusFilterResponse, error)
	GetOrderByOrderID(context.Context, *cookhouse.GetOrderByOrderIDRequest) (*cookhouse.GetOrderByOrderIDResponse, error)
	UpdateOrderStatus(context.Context, *cookhouse.UpdateOrderStatusRequest) (*cookhouse.UpdateOrderStatusResponse, error)
}

type OrderRepositoryManager interface {
	UpsertOrder(context.Context, *data.OrderEntity) (error)
	GetAllOrdersByUserID(context.Context, string) ([]*data.OrderEntity, error)
	GetOrderByOrderID(context.Context, string) (*data.OrderEntity, error)
	GetAllOrdersByStatus(context.Context, *common.OrderStatus) ([]*data.OrderEntity, error)
}