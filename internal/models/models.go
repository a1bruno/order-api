package models

import "github.com/google/uuid"

type Order struct {
	ID       uuid.UUID
	Customer string
	Price    float64
	Refunded bool
}
