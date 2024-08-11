package handlers

import (
	"context"
	"log"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/cookhouse"
	"github.com/Wwreaker007/DIY-menu-service/orders/types"
	"google.golang.org/grpc"
)

type CookHouseManagerHandler struct {
	CookHouseManagerService 	types.CookHouse
	cookhouse.UnimplementedCookHouseServiceServer
}

func NewCookHouseManagerHandler(server *grpc.Server, cms types.CookHouse) {
	cmsHandler := &CookHouseManagerHandler{
		CookHouseManagerService: cms,
	}
	/*
		1. Create service implementation of CookHouseManagerService (cms)
		2. Register CookHouseManagerService in to GRPC server
	*/
	cookhouse.RegisterCookHouseServiceServer(server, cmsHandler)
}

func (h *CookHouseManagerHandler) GetAllOrderByStatusFilter(ctx context.Context, request *cookhouse.GetAllOrderByStatusFilterRequest) (*cookhouse.GetAllOrderByStatusFilterResponse, error) {
	response, err := h.CookHouseManagerService.GetAllOrderByStatusFilter(ctx, request) 
	if err != nil {
		log.Println("Error in fetching all Orders with status filer : " + err.Error())
		return response, err
	}
	return response, err
}

func (h *CookHouseManagerHandler) GetOrderByOrderID(ctx context.Context, request *cookhouse.GetOrderByOrderIDRequest) (*cookhouse.GetOrderByOrderIDResponse, error) {
	response, err := h.CookHouseManagerService.GetOrderByOrderID(ctx, request) 
	if err != nil {
		log.Println("Error in fetching all Orders with OrderID filer : " + err.Error())
		return response, err
	}
	return response, err
}

func (h *CookHouseManagerHandler) UpdateOrderStatus(ctx context.Context, request *cookhouse.UpdateOrderStatusRequest) (*cookhouse.UpdateOrderStatusResponse, error) {
	response, err := h.CookHouseManagerService.UpdateOrderStatus(ctx, request) 
	if err != nil {
		log.Println("Error in updating order status : " + err.Error())
		return response, err
	}
	return response, err
}
