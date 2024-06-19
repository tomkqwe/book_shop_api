package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID       uuid.UUID       `json:"id" db:"id"`
	OrderID  uuid.UUID       `json:"order_id" db:"order_id"`
	BookID   uuid.UUID       `json:"book_id" db:"book_id"`
	Quantity int             `json:"quantity" db:"quantity"`
	Price    decimal.Decimal `json:"price" db:"price"`
}

func (o OrderItem) GetId() uuid.UUID {
	return o.ID
}
