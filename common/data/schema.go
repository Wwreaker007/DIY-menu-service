package data

import (
	"errors"
	"sync"
	"time"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
)

type OrderEntity struct {
	UserID			string				`json:"user_id"`
	Order 			*common.Order		`json:"order"`
	CreatedOn		int64				`json:"created_on"`
	UpdatedOn		int64				`json:"updated_on"`
	CompletedOn		int64				`json:"completed_on,omitempty"`
}

type ThreadSafeOrderEntity struct {
	Orders 		[]*OrderEntity
	mutex		sync.RWMutex
}

func NewThreadSafeOrderEntity() *ThreadSafeOrderEntity {
    return &ThreadSafeOrderEntity{
        Orders: make([]*OrderEntity, 0),
    }
}

func (ts *ThreadSafeOrderEntity) Append(orderEntity *OrderEntity) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	ts.Orders = append(ts.Orders, orderEntity)
}

func (ts *ThreadSafeOrderEntity) GetOrderByOrderID(orderID string) *OrderEntity {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()

	for _, order := range ts.Orders {
		if order.Order.OrderID == &orderID {
			return order
		}
	}
	return nil
}

func (ts *ThreadSafeOrderEntity) GetOrdersByUserID(userID string) []*OrderEntity {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()

	var allOrders []*OrderEntity
	for _, order := range ts.Orders {
		if &order.UserID == &userID {
			allOrders = append(allOrders, order)
		}
	}
	return allOrders
}

func (ts *ThreadSafeOrderEntity) GetOrdersByStatus(status *common.OrderStatus) []*OrderEntity {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()

	var filteredOrders []*OrderEntity
	for _, order := range ts.Orders {
		if order.Order.OrderStatus == status {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}

func (ts *ThreadSafeOrderEntity) UpdateOrder(orderEntity *OrderEntity) error {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	for _, order := range ts.Orders {
		if order.Order.OrderID == orderEntity.Order.OrderID {
			order = orderEntity
			order.UpdatedOn = time.Now().Unix()
			return nil
		}
	}
	return errors.New("No orders with orderID : " + *orderEntity.Order.OrderID)
}