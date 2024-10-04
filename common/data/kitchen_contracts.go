package data

import (
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
)


type CreateOrderRequest struct {
	UserID 		string 				`json:"user_id"`
	Order  		*common.Order  		`json:"order"`
}

type CreateOrderResponse struct {
	Status 		string 				`json:"status"`
	Order  		*common.Order  		`json:"order,omitempty"`
}

type GetOrderRequest struct {
	UserID      string 				`json:"user_id"`
	OrderStatus *common.OrderStatus `json:"order_status,omitempty"`
}

type GetOrderResponse struct {
	Status 		string 				`json:"status"`
	Orders 		[]*common.Order		`json:"orders"`
}

type UpdateOrderRequest struct {
	UserID		string				`json:"user_id"`
	Order		*common.Order		`json:"order"`
}

type UpdateOrderResponse struct {
	Status 		string 				`json:"status"`
	Order		*common.Order		`json:"order"`
}