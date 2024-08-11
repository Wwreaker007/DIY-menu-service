package data

import (
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
)

type GetOrderByStatusFilterRequest struct{
	OrderStatus *common.OrderStatus `json:"order_status"`
}

type GetOrderByStatusFilterResponse struct {
	Status 		string 				``
	Orders		[]*common.Order		`json:"orders"`
}

type GetOrderByOrderIDRequest struct{
	OrderID 	string				`json:"order_id"`
}

type GetOrderByOrderIDResponse struct {
	Status 		string 				`json:"status"`
	Order		*common.Order		`json:"order,omitempty"`
}

type UpdateOrderStatusRequest struct {
	Order		*common.Order		`json:"order"`
}

type UpdateOrderStatusResponse struct {
	Status 		string 				`json:"status"`
	Order		*common.Order		`json:"order"`
}