package cookhousemanager

import (
	"context"
	"log"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/cookhouse"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type CookHouseManagerService struct {
	client cookhouse.CookHouseServiceClient
}

func NewCookhouseManagerService(client cookhouse.CookHouseServiceClient) *CookHouseManagerService {
	return &CookHouseManagerService{
		client:	client,
	}
}

/*
	1. This is used to get all the orders with the order status past as filter to order service.
	2. Return the result order set as response.
*/
func (cms *CookHouseManagerService) GetAllOrdersByStatusFilter(ctx context.Context, request *data.GetOrderByStatusFilterRequest) (*data.GetOrderByStatusFilterResponse, error) {
	grpcRequest := &cookhouse.GetAllOrderByStatusFilterRequest{
		OrderStatus: *request.OrderStatus,
	}

	orders, err := cms.client.GetAllOrderByStatusFilter(ctx, grpcRequest) 
	if err != nil {
		log.Println("error in getting orders: ", err.Error())
		response := &data.GetOrderByStatusFilterResponse{
			Status: "ERROR GETTING ORDERS",
		}
		return response, err
	}
	if orders == nil || len(orders.Orders) == 0 {
		log.Println("no orders with the given status : " + request.OrderStatus.String())
		response := &data.GetOrderByStatusFilterResponse{
			Status: "NO ORDERS FOUND",
		}
		return response, nil
	}

	response := &data.GetOrderByStatusFilterResponse{
		Status: "SUCCESS",
		Orders: orders.Orders,
	}
	return response, nil
}

/*
	1. Fetch order with the matching order ID passed in request to order service.
	2. If no orders found return error response.
	3. Pass the order in response.
*/
func (cms *CookHouseManagerService) GetOrderByOrderID(ctx context.Context, request *data.GetOrderByOrderIDRequest) (*data.GetOrderByOrderIDResponse, error) {
	grpcRequest := &cookhouse.GetOrderByOrderIDRequest{
		OrderID: request.OrderID,
	}

	orders, err := cms.client.GetOrderByOrderID(ctx, grpcRequest) 
	if err != nil {
		log.Println("error in getting order: ", err.Error())
		response := &data.GetOrderByOrderIDResponse{
			Status: "ERROR GETTING ORDER WITH ORDERID : " + request.OrderID,
		}
		return response, err
	}
	if orders == nil {
		log.Println("no order with the given orderID : " + request.OrderID)
		response := &data.GetOrderByOrderIDResponse{
			Status: "NO ORDERS FOUND",
		}
		return response, nil
	}

	response := &data.GetOrderByOrderIDResponse{
		Status: "SUCCESS",
		Order: orders.Order,
	}
	return response, nil
}

/*
	1. Fetch the order with the passed orderID to the order service.
	2. If no orders found return error response.
	3. Update the order with the new order status. 
*/
func (cms *CookHouseManagerService) UpdateOrderStatus(ctx context.Context, request *data.UpdateOrderStatusRequest) (*data.UpdateOrderStatusResponse, error) {
	grpcRequest := &cookhouse.UpdateOrderStatusRequest{
		Order: request.Order,
	}

	orders, err := cms.client.UpdateOrderStatus(ctx, grpcRequest) 
	if err != nil {
		log.Println("error in updating order: ", err.Error())
		response := &data.UpdateOrderStatusResponse{
			Status: "ERROR GETTING ORDER WITH ORDERID : " + *request.Order.OrderID,
		}
		return response, err
	}

	response := &data.UpdateOrderStatusResponse{
		Status: "SUCCESS",
		Order: orders.Order,
	}
	return response, nil
}