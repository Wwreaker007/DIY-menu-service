package utils

import (
	"time"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

func FilterOrdersOnOrderStatus(ordersList []data.OrderEntity, request *orders.GetOrderRequest) (filteredOrders []*common.Order) {
	for _, order := range ordersList {
		if order.Order.OrderStatus == request.OrderStatus {
			filteredOrders = append(filteredOrders, order.Order)
		}
	}
	return filteredOrders
}

func GetNonFilterOrder(ordersList []data.OrderEntity, request *orders.GetOrderRequest) (filteredOrders []*common.Order) {
	for _, order := range ordersList {
		filteredOrders = append(filteredOrders, order.Order)
	}
	return filteredOrders
}

func UpdateOrderStatusInOrderEntity(updatedOrder *common.Order, oldOrderEntity data.OrderEntity) data.OrderEntity {
	if updatedOrder.OrderStatus == common.OrderStatus_ORDER_PLACED.Enum(){
		return data.OrderEntity{
			UserID: oldOrderEntity.UserID,
			Order: updatedOrder,
			CreatedOn: oldOrderEntity.CreatedOn,
			UpdatedOn: time.Now().Unix(),
			CompletedOn: time.Now().Unix(),
		}
	}
	return data.OrderEntity{
		UserID: oldOrderEntity.UserID,
		Order: updatedOrder,
		CreatedOn: oldOrderEntity.CreatedOn,
		UpdatedOn: time.Now().Unix(),
	}
}