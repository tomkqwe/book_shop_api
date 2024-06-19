package http

import (
	"book_shop_api/internal/domain"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BookRequest struct {
	Title         string          `json:"title"`
	Author        string          `json:"author"`
	YearPublished uint            `json:"year_published"`
	Price         decimal.Decimal `json:"price"`
	Category      string          `json:"category"`
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
}

func (r *BookRequest) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("%w: title", domain.ErrRequired)
	}
	if r.Author == "" {
		return fmt.Errorf("%w: author", domain.ErrRequired)
	}
	if r.YearPublished > uint(time.Now().Year()) {
		return fmt.Errorf("%w: year_published", domain.ErrIncorrectYear)
	}
	if r.Price.IsNegative() || r.Price.IsZero() {
		return fmt.Errorf("%w: price", domain.ErrPrice)
	}
	if r.Category == "" {
		return fmt.Errorf("%w: nil category", domain.ErrRequired)
	}
	return nil
}

type BookResponse struct {
	Title         string          `json:"title"`
	Author        string          `json:"author"`
	YearPublushed uint            `json:"year_published"`
	Price         decimal.Decimal `json:"price"`
	Category      string          `json:"category"`
	gorm.Model
}

func (r *BookResponse) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("%w: title", domain.ErrRequired)
	}
	if r.Author == "" {
		return fmt.Errorf("%w: author", domain.ErrRequired)
	}
	if r.YearPublushed > uint(time.Now().Year()) {
		return fmt.Errorf("%w: year_published", domain.ErrIncorrectYear)
	}
	if r.Price.IsNegative() || r.Price.IsZero() {
		return fmt.Errorf("%w: price", domain.ErrPrice)
	}
	if r.Category == "" {
		return fmt.Errorf("%w: nil category", domain.ErrRequired)
	}
	return nil
}
