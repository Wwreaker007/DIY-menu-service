package inmem

import (
	"context"

	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

var imMemoryStore []*data.OrderEntity

type InMemoryDBService struct {

}

func NewInMemoryDbService() *InMemoryDBService {
	return &InMemoryDBService{}
}

func (db *InMemoryDBService) UpsertOrder(ctx context.Context, order *data.OrderEntity) (err error) {

	return err
}

func (db *InMemoryDBService) GetAllOrdersByUserID(ctx context.Context, userID string) (data []*data.OrderEntity, err error) {

	return data, err
}

func (db *InMemoryDBService) GetOrderByOrderID(ctx context.Context, orderID string) (data *data.OrderEntity, err error) {

	return data, err
}