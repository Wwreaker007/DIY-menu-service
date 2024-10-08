package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/Wwreaker007/DIY-menu-service/common/codegen/common"
	"github.com/Wwreaker007/DIY-menu-service/common/data"
)

type PostgresDBService struct{
	client 		*sql.DB
}

func NewPostgresDBService(client *sql.DB) *PostgresDBService{
	return &PostgresDBService{
		client: client,
	}
}

/*
	1. Get the order wrt orderID sent in the request.
	2. Return the result set
*/
func (db *PostgresDBService) GetOrderByOrderID(ctx context.Context, orderID string) (data.OrderEntity, error) {
	query := "SELECT * FROM orders WHERE id = $1"

	var orderEntity data.OrderEntity
	var orderData []byte
	var id int
	row := db.client.QueryRow(query, orderID)
	if row == nil {
		return data.OrderEntity{}, nil
	}
	err := row.Scan(&id, &orderEntity.UserID, &orderData, &orderEntity.CreatedOn, &orderEntity.UpdatedOn, &orderEntity.CompletedOn)
	if err != nil{
		fmt.Println("unable to read data from row : " + err.Error())
		return data.OrderEntity{}, err
	}

	err = json.Unmarshal(orderData, &orderEntity.Order)
	if err != nil {
		fmt.Println("unable to unmarshal order from row : " + err.Error())
		return data.OrderEntity{}, err
	}
	*orderEntity.Order.OrderID = strconv.Itoa(id)
	return orderEntity, nil
}

/*
	1. Fetch the order by orderID.
	2. If found, update the order with the new order.
	3. If not found, insert this as a new entry in the DB.
*/
func (db *PostgresDBService) UpsertOrder(ctx context.Context, orderEntity data.OrderEntity) (error) {
	order, err := db.GetOrderByOrderID(ctx, *orderEntity.Order.OrderID)
	if err != nil {
		fmt.Println("unable to fetch order with orderID : " + *orderEntity.Order.OrderID + ", error : " + err.Error())
		return errors.New("unable to fetch order with orderID : " + *orderEntity.Order.OrderID + ", error : " + err.Error())
	}

	query := "INSERT INTO orders (user_id, order_data) VALUES ($1, $2)"
	if (order == data.OrderEntity{}){
		fmt.Println("no order with orderID : " + *orderEntity.Order.OrderID)
		encodedOrder, err := json.Marshal(orderEntity.Order)
		if err != nil{
			fmt.Println("unable to marshal orderdata into string JSON : " + *orderEntity.Order.OrderID + ", error : " + err.Error())
			return errors.New("unable to marshal orderdata into string JSON : " + *orderEntity.Order.OrderID + ", error : " + err.Error())
		}
		_, err = db.client.Exec(query, orderEntity.UserID, encodedOrder)
		if err != nil{
			fmt.Println("unable to read data from row : " + err.Error())
			return err
		}
		return nil
	}

	// Update query for the same primary key
	updateQuery := "UPDATE orders SET order_data = $1 WHERE id = $2"
	encodedUpdatedOrder, err := json.Marshal(orderEntity.Order)
	if err != nil {
		fmt.Println("unable to marshal orderdata into string JSON : " + *orderEntity.Order.OrderID + ", error : " + err.Error())
		return errors.New("unable to marshal orderdata into string JSON : " + *orderEntity.Order.OrderID + ", error : " + err.Error())
	}
	_, err = db.client.Exec(updateQuery, encodedUpdatedOrder, order.Order.OrderID)
	if err != nil{
		fmt.Println("unable to update data in row : " + *order.Order.OrderID + " error: " + err.Error())
		return err
	}
	return nil
}

/*
	1. Get all the the orders wrt userIDs
	2. Return the result set
*/
func (db *PostgresDBService) GetAllOrdersByUserID(ctx context.Context, userID string) ([]data.OrderEntity, error) {
	query := "SELECT * FROM orders WHERE user_id = $1"

	var orderEntity data.OrderEntity
	rows, err := db.client.Query(query, userID)
	if err != nil {
		fmt.Println("unable to fetch orders for userid : " + userID + " error : " + err.Error())
		return nil, errors.New("unable to fetch orders for userid : " + userID + " error : " + err.Error())
	}

	var orders []data.OrderEntity
	for rows.Next() {
		var order data.OrderEntity
		var encodedOrder []byte
		var id int
		err := rows.Scan(&id, &orderEntity.UserID, &encodedOrder, &orderEntity.CreatedOn, &orderEntity.UpdatedOn, &orderEntity.CompletedOn)
		if err != nil{
			fmt.Println("unable to read data from row : " + err.Error())
			return nil, err
		}
		err = json.Unmarshal(encodedOrder, &order.Order)
		if err != nil {
			fmt.Println("unable to unmarshal order from row : " + err.Error())
			return nil, err
		}
		*orderEntity.Order.OrderID = strconv.Itoa(id)
		orders = append(orders, order)
	}
	return orders, nil
}


/*
	1. Get all the the orders wrt orderStatus as filter.
	2. Return the result set.
*/
func (db *PostgresDBService) GetAllOrdersByStatus(ctx context.Context, orderStatus common.OrderStatus) ([]data.OrderEntity, error) {
	query := "SELECT * FROM orders"

	var orderEntity data.OrderEntity
	rows, err := db.client.Query(query)
	if err != nil {
		fmt.Println("unable to fetch orders :,  error : " + err.Error())
		return nil, errors.New("unable to fetch orders, error : " + err.Error())
	}

	var orders []data.OrderEntity
	for rows.Next() {
		var order data.OrderEntity
		var encodedOrder []byte
		var id int

		err := rows.Scan(&id, &orderEntity.UserID, &encodedOrder, &orderEntity.CreatedOn, &orderEntity.UpdatedOn, &orderEntity.CompletedOn)
		if err != nil{
			fmt.Println("unable to read data from row : " + err.Error())
			return nil, err
		}
		err = json.Unmarshal(encodedOrder, &order.Order)
		if err != nil {
			fmt.Println("unable to unmarshal order from row : " + err.Error())
			return nil, err
		}
		*orderEntity.Order.OrderID = strconv.Itoa(id)
		if(orderStatus == *order.Order.OrderStatus){
			orders = append(orders, order)
		}
	}
	return orders, nil
}