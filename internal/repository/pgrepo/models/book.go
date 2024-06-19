package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Book struct {
	Id            uuid.UUID       `json:"id" db:"id"`
	Title         string          `json:"title" db:"title"`
	Author        string          `json:"author" db:"author"`
	YearPublished uint            `json:"year_published" db:"year_published"`
	Price         decimal.Decimal `json:"price" db:"price"`
	Category      string          `json:"category" db:"category"`
}

func (s Book) GetId() uuid.UUID {
	return s.Id
}
