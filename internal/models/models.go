package models

import "github.com/google/uuid"

type Order struct {
	ID       uuid.UUID `json:"user_id"`
	Customer string    `json:"customer_name"`
	Price    float64   `json:"product_price"`
	Refunded bool      `json:"refunded"`
}

type CreateOrder struct { //essa struct serve para implementar a lógica do Add() sem parametros extras do usuario
	Customer string  `json:"customer"`
	Price    float64 `json:"price"`
}

type ResponseOrder struct {
	NewOrderID uuid.UUID `json:"new_order_id"` //essa struct vai ser usada para retornar o id quando o usuario adicionar um novo pedido
}
