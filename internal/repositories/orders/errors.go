package orders

import "errors"

var (
	ErrNotFound             = errors.New("Pedido não encontrado")
	ErrOrderAlreadyRefunded = errors.New("Pedido já foi estornado")
)
