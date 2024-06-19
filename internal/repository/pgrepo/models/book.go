package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Эта модель нужна для общения с БД
type Book struct {
	Title         string
	Author        string
	YearPublished uint
	Price         decimal.Decimal
	Category      string
	Id            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	b.Id = uuid.New()
	b.CreatedAt = time.Now()
	return nil
}
