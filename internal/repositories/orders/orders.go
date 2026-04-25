package orders

import (
	"database/sql"

	"github.com/a1bruno/order-api/internal/models"
	"github.com/google/uuid"
)

type Orders struct {
	connection *sql.DB
}

func New(connection *sql.DB) *Orders {
	return &Orders{connection: connection}
}

func (o *Orders) GetAll() ([]models.Order, error) {
	rows, err := o.connection.Query("SELECT id, customer, price, refunded FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Customer, &order.Price, &order.Refunded)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (o *Orders) Add(newOrder models.Order) error {
	_, err := o.connection.Exec(
		"INSERT INTO orders (id, customer, price, refunded) VALUES ($1, $2, $3, $4)",
		newOrder.ID, newOrder.Customer, newOrder.Price, newOrder.Refunded,
	)
	return err
}

func (o *Orders) Refund(id uuid.UUID) error {
	result, err := o.connection.Exec(
		"UPDATE orders SET refunded = TRUE WHERE id = $1 AND refunded = FALSE",
		id,
	)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}
