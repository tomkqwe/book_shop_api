package models

import (
	"github.com/google/uuid"
)

type CartItem struct {
	ID     uuid.UUID `json:"id" db:"id"`
	CartID uuid.UUID `json:"cart_id" db:"cart_id"`
	BookID uuid.UUID `json:"book_id" db:"book_id"`
	Count  uint      `json:"count" db:"count"`
}

func (c CartItem) GetId() uuid.UUID {
	return c.ID
}
