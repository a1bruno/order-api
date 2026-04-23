package orders

import (
	"github.com/a1bruno/order-api/internal/models"
	"github.com/google/uuid"
)

type Orders struct {
	orders []models.Order
}

func New() *Orders {
	return &Orders{}
}

func (o *Orders) GetAll() ([]models.Order, error) {
	return o.orders, nil
}

func (o *Orders) Add(newOrder models.Order) error {
	o.orders = append(o.orders, newOrder)
	return nil
}

func (o *Orders) Refund(id uuid.UUID) error {
	for i, order := range o.orders {
		if order.ID == id {
			if order.Refunded {
				return ErrOrderAlreadyRefunded
			}
			o.orders[i].Refunded = true
			return nil
		}
	}
	return ErrNotFound
}
