package data

import (
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
)

type CreateOrderRequest struct {
	UserID 		string 				`json:"user_id"`
	Order  		*orders.Order  		`json:"order"`
}

type CreateOrderResponse struct {
	Status 		string 				`json:"status"`
	Order  		*orders.Order  		`json:"order,omitempty"`
}

type GetOrderRequest struct {
	UserID      string 				`json:"user_id"`
	OrderStatus *orders.OrderStatus `json:"order_status,omitempty"`
}

type GetOrderResponse struct {
	Status 		string 				`json:"user_id"`
	Orders 		[]*orders.Order		`json:"orders"`
}

type GetOrderByOrderIDRequest struct{
	UserID 		string				`json:"user_id"`
	Order		*orders.Order		`json:"order"`
}

type GetOrderByOrderIDResponse struct {
	Status 		string 				`json:"status"`
	Order		*orders.Order		`json:"order,omitempty"`
}

type UpdateOrderRequest struct {
	UserID		string				`json:"user_id"`
	Order		*orders.Order		`json:"order"`
}

type UpdateOrderResponse struct {
	Status 		string 				`json:"status"`
	Order		*orders.Order		`json:"order"`
}