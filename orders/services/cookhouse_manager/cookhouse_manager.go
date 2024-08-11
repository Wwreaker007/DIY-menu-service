package cookhousemanager

import (
	"context"
	"log"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
	"github.com/Wwreaker007/DIY-menu-service/common/codegen/cookhouse"
	"github.com/Wwreaker007/DIY-menu-service/orders/types"
	"github.com/Wwreaker007/DIY-menu-service/orders/utils"
)

type CookHouseManagerService struct {
	db types.OrderRepositoryManager
}

func NewCookHouseManagerService(db types.OrderRepositoryManager) *CookHouseManagerService {
	return &CookHouseManagerService{
		db: db,
	}
}

func (cms *CookHouseManagerService) GetAllOrderByStatusFilter(ctx context.Context, request *cookhouse.GetAllOrderByStatusFilterRequest) (*cookhouse.GetAllOrderByStatusFilterResponse, error) {
	orders, err := cms.db.GetAllOrdersByStatus(ctx, &request.OrderStatus)
	if err != nil {
		log.Println("error in fetching orders: ", err.Error())
		response := &cookhouse.GetAllOrderByStatusFilterResponse{
			Status: "ERROR",
		}
		return response, err
	}
	if len(orders) == 0 || orders == nil {
		log.Println("No orders present with the status: " + request.OrderStatus.String())
		response := &cookhouse.GetAllOrderByStatusFilterResponse{
			Status: "NO OREDERS WITH STATUS : " + request.OrderStatus.String(),
		}
		return response, nil
	}

	var transformedOrders []*common.Order
	for _, order := range orders {
		transformedOrders = append(transformedOrders, order.Order)
	}
	response := &cookhouse.GetAllOrderByStatusFilterResponse{
		Status: "SUCCESS",
		Orders: transformedOrders,
	}
	return response, nil
}

func (cms *CookHouseManagerService) GetOrderByOrderID(ctx context.Context, request *cookhouse.GetOrderByOrderIDRequest) (*cookhouse.GetOrderByOrderIDResponse, error) {
	order, err := cms.db.GetOrderByOrderID(ctx, request.OrderID)
	if err != nil {
		log.Println("error in fetching order bu orderID: ", err.Error())
		response := &cookhouse.GetOrderByOrderIDResponse{
			Status: "ERROR",
		}
		return response, err
	}
	if order == nil {
		log.Println("No orders present with orderID : " + request.OrderID)
		response := &cookhouse.GetOrderByOrderIDResponse{
			Status: "NO OREDERS WITH ORDERID : " + request.OrderID,
		}
		return response, nil
	}

	response := &cookhouse.GetOrderByOrderIDResponse{
		Status: "SUCCESS",
		Order: order.Order,
	}
	return response, nil
}

func (cms *CookHouseManagerService) UpdateOrderStatus(ctx context.Context, request *cookhouse.UpdateOrderStatusRequest) (*cookhouse.UpdateOrderStatusResponse, error) {
	order, err := cms.db.GetOrderByOrderID(ctx, *request.Order.OrderID)
	if err != nil {
		log.Println("error in fetching order by orderID: ", err.Error())
		response := &cookhouse.UpdateOrderStatusResponse{
			Status: "ERROR",
		}
		return response, err
	}
	if order == nil {
		log.Println("No orders present with orderID : " + *request.Order.OrderID)
		response := &cookhouse.UpdateOrderStatusResponse{
			Status: "NO OREDERS WITH ORDERID : " + *request.Order.OrderID,
		}
		return response, nil
	}

	updatedOrderEntity := utils.UpdateOrderStatusInOrderEntity(request.Order, order)
	err = cms.db.UpsertOrder(ctx, updatedOrderEntity)
	if err != nil {
		log.Println("ERROR IN UPDATING ORDER STATUS : " + *request.Order.OrderID)
		response := &cookhouse.UpdateOrderStatusResponse{
			Status: "UNABLE TO UPDATE ORDER STATUS " + *request.Order.OrderID,
		}
		return response, nil
	}

	response := &cookhouse.UpdateOrderStatusResponse{
		Status: "SUCCESS",
		Order: request.Order,
	}
	return response, nil
}