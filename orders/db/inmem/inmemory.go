package inmem

import (
	"context"
	"time"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type InMemoryDBService struct {
	inMemoryStore 	[]*data.OrderEntity
}

func NewInMemoryDbService(inMemoryStore []*data.OrderEntity) *InMemoryDBService {
	return &InMemoryDBService{
		inMemoryStore: inMemoryStore,
	}
}

/*
	1. Perform linear search in the datastore to find an existing order with passed orderID
	2. If found, update it.
	3. If not found, append this in the inMemoryStore.
*/
func (db *InMemoryDBService) UpsertOrder(ctx context.Context, order *data.OrderEntity) (err error) {
	// Check for an existing orde with the passed OrderID
	// Update and return no error.
	for _, orderEntity := range db.inMemoryStore {
		if orderEntity.Order.OrderID == order.Order.OrderID {
			orderEntity.Order = order.Order
			orderEntity.UpdatedOn = time.Now().Unix()
			return nil
		}
	}

	// If not found, append the new order at the end of the inMemoryStore
	updatedDataStore := append(db.inMemoryStore, order)
	db.inMemoryStore = updatedDataStore
	return err
}

/*
	1. Get all the the orders wrt userIDs
	2. Return the result set
*/
func (db *InMemoryDBService) GetAllOrdersByUserID(ctx context.Context, userID string) ([]*data.OrderEntity, error) {
	// Perform linear search on the inMemoryStore and collect the orders wrt userID
	var filteredOrders []*data.OrderEntity
	for _, orderEntity := range db.inMemoryStore {
		if orderEntity.UserID == userID {
			filteredOrders = append(filteredOrders, orderEntity)
		}
	}
	return filteredOrders, nil
}

/*
	1. Get all the the orders wrt orderID
	2. Return the result set
*/
func (db *InMemoryDBService) GetOrderByOrderID(ctx context.Context, orderID string) (*data.OrderEntity, error) {
	// Perform linear search on the inMemoryStore and collect the orders wrt userID
	var result *data.OrderEntity
	for _, orderEntity := range db.inMemoryStore {
		if orderEntity.Order.OrderID == &orderID {
			result = orderEntity
			break
		}
	}
	return result, nil
}

/*
	1. Get all the the orders wrt orderStatus as filter.
	2. Return the result set.
*/
func (db *InMemoryDBService) GetAllOrdersByStatus(ctx context.Context, orderStatus *common.OrderStatus) ([]*data.OrderEntity, error) {
	// Perform linear search on the inMemoryStore and collect the orders wrt orderStatus
	var filteredOrders []*data.OrderEntity
	for _, orderEntity := range db.inMemoryStore {
		if orderEntity.Order.OrderStatus == orderStatus {
			filteredOrders = append(filteredOrders, orderEntity)
		}
	}
	return filteredOrders, nil
}