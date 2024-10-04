package inmem

import (
	"context"
	"errors"
	"time"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type InMemoryDBService struct {
	inMemoryStore 	*data.ThreadSafeOrderEntity
}

func NewInMemoryDbService(inMemoryStore *data.ThreadSafeOrderEntity) *InMemoryDBService {
	return &InMemoryDBService{
		inMemoryStore: inMemoryStore,
	}
}

/*
	1. Perform linear search in the datastore to find an existing order with passed orderID
	2. If found, update it.
	3. If not found, append this in the inMemoryStore.
*/
func (db *InMemoryDBService) UpsertOrder(ctx context.Context, order data.OrderEntity) (err error) {
	// Check for an existing order with the passed OrderID
	// Update and return no error.
	existingOrder := db.inMemoryStore.GetOrderByOrderID(*order.Order.OrderID)
	if (existingOrder != data.OrderEntity{}) {
		if order.Order.OrderStatus == common.OrderStatus_ORDER_READY.Enum() {
			order.CompletedOn = time.Now().Unix()
		}
		err = db.inMemoryStore.UpdateOrder(order)
		return err
	}

	// If not found, append the new order at the end of the inMemoryStore
	db.inMemoryStore.Append(order)
	return err
}

/*
	1. Get all the the orders wrt userIDs
	2. Return the result set
*/
func (db *InMemoryDBService) GetAllOrdersByUserID(ctx context.Context, userID string) ([]data.OrderEntity, error) {
	// Perform linear search on the inMemoryStore and collect the orders wrt userID
	allOrders := db.inMemoryStore.GetOrdersByUserID(userID)
	if len(allOrders) == 0 {
		return nil, errors.New("No orders for userID : " + userID)
	}
	return allOrders, nil
}

/*
	1. Get all the the orders wrt orderID
	2. Return the result set
*/
func (db *InMemoryDBService) GetOrderByOrderID(ctx context.Context, orderID string) (data.OrderEntity, error) {
	// Perform linear search on the inMemoryStore and collect the orders wrt userID
	order := db.inMemoryStore.GetOrderByOrderID(orderID)
	if (order == data.OrderEntity{}) {
		return order, errors.New("No order with OrderID : " + orderID)
	}
	return order, nil
}

/*
	1. Get all the the orders wrt orderStatus as filter.
	2. Return the result set.
*/
func (db *InMemoryDBService) GetAllOrdersByStatus(ctx context.Context, orderStatus common.OrderStatus) ([]data.OrderEntity, error) {
	// Perform linear search on the inMemoryStore and collect the orders wrt orderStatus
	filteredOrders := db.inMemoryStore.GetOrdersByStatus(orderStatus)
	if len(filteredOrders) == 0 {
		return nil, errors.New("No orders with status : " + orderStatus.String())
	}
	return filteredOrders, nil
}