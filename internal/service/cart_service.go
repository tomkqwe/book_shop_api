package service

import (
	"book_shop_api/internal/repository/pgrepo"
	"book_shop_api/internal/repository/pgrepo/models"
	"context"

	"github.com/google/uuid"
)

type cartService struct {
	r pgrepo.CRUDRepository[models.Cart]
}

func NewCartService(r pgrepo.CRUDRepository[models.Cart]) Service[models.Cart] {
	return cartService{r: r}
}

func (s cartService) GetEntity(ctx context.Context, id uuid.UUID) (models.Cart, error) {
	return s.r.GetEntity(ctx, id)
}

func (s cartService) PostEntity(ctx context.Context, entity models.Cart) (models.Cart, error) {
	return s.r.PostEntity(ctx, entity)
}

func (s cartService) GetEntities(ctx context.Context) ([]models.Cart, error) {
	return s.r.GetEntities(ctx)
}

func (s cartService) PutEntity(ctx context.Context, id uuid.UUID, newEntity models.Cart) (bool, error) {
	return s.r.PutEntity(ctx, id, newEntity)
}

func (s cartService) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.r.DeleteEntity(ctx, id)
}
