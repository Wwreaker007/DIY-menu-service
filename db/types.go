package db

import (
	"context"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type OrderRepositoryManager interface {
	UpsertOrder(context.Context, *data.OrderEntity) (error)
	GetAllOrdersByUserID(context.Context, string) ([]*data.OrderEntity, error)
	GetOrderByOrderID(context.Context, string) (*data.OrderEntity, error)
}