package service

import (
	"book_shop_api/internal/repository/pgrepo"
	"book_shop_api/internal/repository/pgrepo/models"
	"context"

	"github.com/google/uuid"
)

type cartItemService struct {
	r pgrepo.CRUDRepository[models.CartItem]
}

func NewCartItemService(r pgrepo.CRUDRepository[models.CartItem]) Service[models.CartItem] {
	return cartItemService{r: r}
}

func (s cartItemService) GetEntity(ctx context.Context, id uuid.UUID) (models.CartItem, error) {
	return s.r.GetEntity(ctx, id)
}

func (s cartItemService) PostEntity(ctx context.Context, entity models.CartItem) (models.CartItem, error) {
	return s.r.PostEntity(ctx, entity)
}

func (s cartItemService) GetEntities(ctx context.Context) ([]models.CartItem, error) {
	return s.r.GetEntities(ctx)
}

func (s cartItemService) PutEntity(ctx context.Context, id uuid.UUID, newEntity models.CartItem) (bool, error) {
	return s.r.PutEntity(ctx, id, newEntity)
}

func (s cartItemService) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.r.DeleteEntity(ctx, id)
}
