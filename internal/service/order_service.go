package service

import (
	"book_shop_api/internal/repository/pgrepo"
	"book_shop_api/internal/repository/pgrepo/models"
	"context"

	"github.com/google/uuid"
)

type orderService struct {
	r pgrepo.CRUDRepository[models.Order]
}

func NewOrderService(r pgrepo.CRUDRepository[models.Order]) Service[models.Order] {
	return orderService{r: r}
}

func (s orderService) GetEntity(ctx context.Context, id uuid.UUID) (models.Order, error) {
	return s.r.GetEntity(ctx, id)
}

func (s orderService) PostEntity(ctx context.Context, entity models.Order) (models.Order, error) {
	return s.r.PostEntity(ctx, entity)
}

func (s orderService) GetEntities(ctx context.Context) ([]models.Order, error) {
	return s.r.GetEntities(ctx)
}

func (s orderService) PutEntity(ctx context.Context, id uuid.UUID, newEntity models.Order) (bool, error) {
	return s.r.PutEntity(ctx, id, newEntity)
}

func (s orderService) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.r.DeleteEntity(ctx, id)
}
