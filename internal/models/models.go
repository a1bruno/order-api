package models

import "github.com/google/uuid"

type Order struct {
	ID       uuid.UUID
	Customer string
	Price    float64
	Refunded bool
}

type CreateOrder struct { //essa struct serve para implementar a lógica do Add() sem parametros extras do usuario
	Customer string
	Price    float64
}

type ResponseOrder struct {
	NewOrderID uuid.UUID `json:"new_order_id"` //essa struct vai ser usada para retornar o id quando o usuario adicionar um novo pedido
}
