package service

import (
	"book_shop_api/internal/repository/pgrepo"
	"book_shop_api/internal/repository/pgrepo/models"
	"context"

	"github.com/google/uuid"
)

type bookService struct {
	r pgrepo.CRUDRepository[models.Book]
}

func NewBookService(r pgrepo.CRUDRepository[models.Book]) Service[models.Book] {
	return bookService{r: r}
}

func (s bookService) GetEntity(ctx context.Context, id uuid.UUID) (models.Book, error) {
	return s.r.GetEntity(ctx, id)
}

func (s bookService) PostEntity(ctx context.Context, entity models.Book) (models.Book, error) {
	return s.r.PostEntity(ctx, entity)
}

func (s bookService) GetEntities(ctx context.Context) ([]models.Book, error) {
	return s.r.GetEntities(ctx)
}

func (s bookService) PutEntity(ctx context.Context, id uuid.UUID, newEntity models.Book) (bool, error) {
	return s.r.PutEntity(ctx, id, newEntity)
}

func (s bookService) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.r.DeleteEntity(ctx, id)
}
