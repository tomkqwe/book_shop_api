package service

import (
	"book_shop_api/internal/repository/pgrepo"
	"book_shop_api/internal/repository/pgrepo/models"
	"context"

	"github.com/google/uuid"
)

type orderItemService struct {
	r pgrepo.CRUDRepository[models.OrderItem]
}

func NewOrderItemService(r pgrepo.CRUDRepository[models.OrderItem]) Service[models.OrderItem] {
	return orderItemService{r: r}
}

func (s orderItemService) GetEntity(ctx context.Context, id uuid.UUID) (models.OrderItem, error) {
	return s.r.GetEntity(ctx, id)
}

func (s orderItemService) PostEntity(ctx context.Context, entity models.OrderItem) (models.OrderItem, error) {
	return s.r.PostEntity(ctx, entity)
}

func (s orderItemService) GetEntities(ctx context.Context) ([]models.OrderItem, error) {
	return s.r.GetEntities(ctx)
}

func (s orderItemService) PutEntity(ctx context.Context, id uuid.UUID, newEntity models.OrderItem) (bool, error) {
	return s.r.PutEntity(ctx, id, newEntity)
}

func (s orderItemService) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.r.DeleteEntity(ctx, id)
}
