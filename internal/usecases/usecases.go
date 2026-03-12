package usecases

import (
	"github.com/a1bruno/order-api/internal/models"
	"github.com/a1bruno/order-api/internal/repositories"
	"github.com/google/uuid"
)

type UseCases struct {
	repos repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		repos: *repos,
	}
}

func (u *UseCases) GetAll() []models.Order {
	orders := u.repos.Order.GetAll()

	return orders
}

func (u *UseCases) Add(newOrder models.CreateOrder) uuid.UUID {
	repoRequest := models.Order{
		ID:       uuid.New(),
		Customer: newOrder.Customer,
		Price:    newOrder.Price,
		Refunded: false,
	}

	u.repos.Order.Add(repoRequest)
	return repoRequest.ID
}

func (u *UseCases) Refund(id uuid.UUID) {
	u.repos.Order.Refund(id)
}
