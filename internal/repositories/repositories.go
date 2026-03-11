package repositories

import (
	"github.com/a1bruno/order-api/internal/models"
	"github.com/google/uuid"
)

type Repositories struct {
	Order interface {
		GetAll() []models.Order
		Add(newUser models.Order)
		Refund(id uuid.UUID)
	}
}

func New() *Repositories {
	return &Repositories{}
}
