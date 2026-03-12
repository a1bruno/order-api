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
