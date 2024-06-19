package service

import (
	"context"

	"github.com/google/uuid"
)

type Entity interface {
	GetId() uuid.UUID
}

type Service[T Entity] interface {
	GetEntity(ctx context.Context, id uuid.UUID) (T, error)
	PostEntity(ctx context.Context, entity T) (T, error)
	GetEntities(ctx context.Context) ([]T, error)
	PutEntity(ctx context.Context, id uuid.UUID, newEntity T) (bool, error)
	DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error)
}
