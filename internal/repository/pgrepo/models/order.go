package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID          uuid.UUID       `json:"id" db:"id"`
	UserID      uuid.UUID       `json:"user_id" db:"user_id"`
	OrderDate   time.Time       `json:"order_date" db:"order_date"`
	TotalAmount decimal.Decimal `json:"total_amount" db:"total_amount"`
}

func (o Order) GetId() uuid.UUID {
	return o.ID
}
