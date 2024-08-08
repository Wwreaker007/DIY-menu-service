package utils

import (
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

func FilterOrdersOnOrderStatus(ordersList []*data.OrderEntity, request *orders.GetOrderRequest) (filteredOrders []*orders.Order) {
	for _, order := range ordersList {
		if order.Order.OrderStatus == request.OrderStatus {
			filteredOrders = append(filteredOrders, order.Order)
		}
	}
	return filteredOrders
}