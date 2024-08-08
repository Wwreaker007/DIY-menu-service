package data

import "github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"

type OrderEntity struct {
	UserID			string				`json:"user_id"`
	Order 			*orders.Order		`json:"order"`
	CreatedOn		int64				`json:"created_on"`
	UpdatedOn		int64				`json:"updated_on"`
	CompletedOn		int64				`json:"completed_on,omitempty"`
}