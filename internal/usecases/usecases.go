package usecases

import (
	"fmt"

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

func (u *UseCases) GetAll() ([]models.Order, error) {
	orders, err := u.repos.Order.GetAll()
	if err != nil {
		return nil, fmt.Errorf("falha ao buscar pedidos: %w", err)
	}
	return orders, nil
}

func (u *UseCases) Add(newOrder models.CreateOrder) (uuid.UUID, error) {
	repoRequest := models.Order{
		ID:       uuid.New(),
		Customer: newOrder.Customer,
		Price:    newOrder.Price,
		Refunded: false,
	}

	err := u.repos.Order.Add(repoRequest)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Erro ao adicionar o pedido %w", err)
	}

	return repoRequest.ID, nil
}

func (u *UseCases) Refund(id uuid.UUID) error {
	err := u.repos.Order.Refund(id)
	if err != nil {
		return fmt.Errorf("falha ao reembolsar pedido %w", err)
	}
	return nil
}
