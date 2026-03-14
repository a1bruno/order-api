package orders

import (
	"github.com/a1bruno/order-api/internal/models"
	"github.com/google/uuid"
)

type Orders struct {
	orders []models.Order //definindo order como um slice para o armazenamento de maneira local;
}

func New() *Orders {
	return &Orders{
		orders: []models.Order{
			{
				ID:       uuid.New(),
				Customer: "Eddard Stark",
				Price:    8.90,
				Refunded: false,
			},
			{
				ID:       uuid.New(),
				Customer: "Robert Baratheon",
				Price:    14.30,
				Refunded: false,
			},
		},
	} //inicializando o orders como um slice vazio;
}

func (o *Orders) GetAll() []models.Order {
	return o.orders
}

func (o *Orders) Add(newOrder models.Order) {
	o.orders = append(o.orders, newOrder)
}

func (o *Orders) Refund(id uuid.UUID) {
	for i, order := range o.orders {
		if order.ID == id {
			o.orders[i].Refunded = true
			return
		}
	}
}
