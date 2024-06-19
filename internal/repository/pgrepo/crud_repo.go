package pgrepo

import (
	"context"

	"github.com/google/uuid"
)

type CRUDRepository[T any] interface {
	PostEntity(ctx context.Context, model T) (T, error)
	GetEntity(ctx context.Context, id uuid.UUID) (T, error)
	GetEntities(ctx context.Context) ([]T, error)
	PutEntity(ctx context.Context, id uuid.UUID, newModel T) (bool, error)
	DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error)
}
