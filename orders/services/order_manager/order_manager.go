package ordermanager

import (
	"context"
	"log"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
	"github.com/Wwreaker007/DIY-menu-service/orders/types"
	"github.com/Wwreaker007/DIY-menu-service/orders/utils"
)

type OrderManagerService struct {
	db 		types.OrderRepositoryManager
}

func NewOrderManagerService(db types.OrderRepositoryManager) *OrderManagerService {
	return &OrderManagerService{
		db: db,
	}
}

/*
	1. Since the order is created, chaneg the OrderStatus to ORDER_PLACED
	2. Store this new order instance in the DB
	3. Return the order as response with the updated details in the response.
*/
func (oms *OrderManagerService) CreateOrder(ctx context.Context, request *orders.CreateOrderRequest) (response *orders.CreateOrderResponse, err error) {
	// Update the order status to ORDER_PLACED
	orderStatus := common.OrderStatus_ORDER_PLACED
	request.Order.OrderStatus = &orderStatus

	// Create a new entity and store it in the DB
	newOrderEntity := data.OrderEntity{
		UserID: request.UserID,
		Order: request.Order,
	}
	err = oms.db.UpsertOrder(ctx, newOrderEntity)
	if err != nil {
		log.Println("error in adding order to DB: ", err.Error())
		response = &orders.CreateOrderResponse {
			Status: "ERROR : " + err.Error(),
		}
		return response, err
	}

	// Send the updated order in the response
	response = &orders.CreateOrderResponse{
		Status: "SUCCESS",
		Order: newOrderEntity.Order,
	}
	return response, err
}

/*
	1. Get all the orders wrt userIDs.
	2. Check for the filter of ORDER_STATUS (If it is populated, filter the orders with ORDER_STATUS, else return all orders)
	3. Send these back as response.
*/
func (oms *OrderManagerService) GetOrder(ctx context.Context, request *orders.GetOrderRequest) (response *orders.GetOrderResponse, err error) {
	fetchedOrders, err := oms.db.GetAllOrdersByUserID(ctx, request.UserID)
	if err != nil {
		log.Println("error in fetching orders :", err.Error())
		response = &orders.GetOrderResponse{
			Status: "ERROR: " + err.Error(),
		}
		return response, err
	}
	if len(fetchedOrders) == 0 || fetchedOrders == nil{
		log.Println("no orders found for this userID :", request.UserID)
		response = &orders.GetOrderResponse{
			Status: "NO ORDERS FOUND !",
		}
		return response, err
	}

	// Check if ORDER_STATUS filter is to be applied, else return all orders
	var filteredOrders []*common.Order
	if(request.OrderStatus != nil) {
		filteredOrders = utils.FilterOrdersOnOrderStatus(fetchedOrders, request)
	}else {
		filteredOrders = utils.GetNonFilterOrder(fetchedOrders, request)
	}

	
	response = &orders.GetOrderResponse {
		Status: "SUCCESS",
		Orders: filteredOrders,
	}
	return response, err
}

/*
	1. Fetch the order for the specific orderID.
	2. Update the order with the new data.
	3. Add the updated order in the DB.
*/
func (oms *OrderManagerService) UpdateOrder(ctx context.Context, request *orders.UpdateOrderRequest) (response *orders.UpdateOrderResponse, err error) {
	// Fetch the order with the orderID and perform checks on the result
	order, err := oms.db.GetOrderByOrderID(ctx, *request.Order.OrderID)
	if err != nil {
		log.Println("error in getting the order : ", err.Error())
		response = &orders.UpdateOrderResponse{
			Status: "ERROR",
		}
		return response, err
	}
	if (order == data.OrderEntity{}) {
		log.Println("no order found with orderID : " + *request.Order.OrderID)
		response = &orders.UpdateOrderResponse{
			Status: "NO ORDER FOUND",
		}
		return response, err
	}

	// Update the order with the updated order details
	updatedOrderEntity := data.OrderEntity{
		UserID: request.UserID,
		Order: request.Order,
	}

	// Add the updated order in the DB
	err = oms.db.UpsertOrder(ctx, updatedOrderEntity)
	if err != nil {
		log.Println("unable to update the order with OrderID : " + *request.Order.OrderID)
		response = &orders.UpdateOrderResponse{
			Status: "ORDER UPDATE FAILED",
		}
		return response, err
	}

	response = &orders.UpdateOrderResponse{
		Status: "UPDATED",
		Order: updatedOrderEntity.Order,
	}
	return response, err
}