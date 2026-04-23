package repositories

import (
	"database/sql"

	"github.com/a1bruno/order-api/internal/models"
	"github.com/a1bruno/order-api/internal/repositories/orders"
	"github.com/google/uuid"
)

type Repositories struct {
	Order interface {
		GetAll() ([]models.Order, error)
		Add(newUser models.Order) error
		Refund(id uuid.UUID) error
	}
}

func New(connection *sql.DB) *Repositories {
	return &Repositories{
		Order: orders.New(connection),
	}
}
