package kitchenmanager

import (
	"context"
	"fmt"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/orders"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
	"github.com/Wwreaker007/DIY-menu-service/kitchen/utils"
)

type KitchenManagerService struct{
	client 	orders.OrderServiceClient
}

func NewKitchenManagerService(client orders.OrderServiceClient) *KitchenManagerService {
	return &KitchenManagerService{
		client: client,
	}
}

func (kms *KitchenManagerService) CreateOrder(ctx context.Context, request *data.CreateOrderRequest) (*data.CreateOrderResponse, error) {
	grpcRequest := utils.TransformToGRPCCreateOrderRequest(request)
	response, err := kms.client.CreateOrder(ctx, grpcRequest)
	if err != nil {
		fmt.Println("error in CreateOrder RPC", err.Error())
		return nil, err
	}
	return utils.TransformToHTTPCreateOrderResponse(response), nil
}

func (kms *KitchenManagerService) GetOrder(ctx context.Context, request *data.GetOrderRequest) (*data.GetOrderResponse, error) {
	grpcRequest := utils.TransformToGRPCGetOrderRequest(request)
	response, err := kms.client.GetOrder(ctx, grpcRequest)
	if err != nil {
		fmt.Println("error in GetOrder RPC", err.Error())
		return nil, err
	}
	return utils.TransformToHTTPGetOrderRespone(response), nil
}

func (kms *KitchenManagerService) UpdateOrder(ctx context.Context, request *data.UpdateOrderRequest) (*data.UpdateOrderResponse, error) {
	grpcRequest := utils.TransformToGRPCUpdateOrderRequest(request)
	response, err := kms.client.UpdateOrder(ctx, grpcRequest)
	if err != nil {
		fmt.Println("error in UpdateOrder RPC", err.Error())
		return nil, err
	}
	return utils.TransformToHTTPUpdateOrderResponse(response), nil
}