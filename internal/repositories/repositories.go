package repositories

import (
	"github.com/a1bruno/order-api/internal/models"
	"github.com/a1bruno/order-api/internal/repositories/orders"
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
	return &Repositories{
		Order: orders.New(), //inicialização do armazenamento de orders, não implementação de método*
	}
}
