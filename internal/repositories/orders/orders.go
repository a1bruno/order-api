package orders

import (
	"github.com/a1bruno/order-api/internal/models"
	"github.com/google/uuid"
)

type Orders struct {
	users []models.Order
}

func New() *Orders {
	return &Orders{users: make([]models.Order, 0)}
}

func (o Orders) GetAll() []models.Order

func (o Orders) Add(newOrder models.Order)

func (o Orders) Refund(id uuid.UUID)
