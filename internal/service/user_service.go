package service

import (
	"book_shop_api/internal/repository/pgrepo"
	"book_shop_api/internal/repository/pgrepo/models"
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	pgrepo.CRUDRepository[models.User]
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type userService struct {
	r pgrepo.UserRepository
}

func NewUserService(r pgrepo.UserRepository) UserService {
	return userService{r: r}
}

func (s userService) GetEntity(ctx context.Context, id uuid.UUID) (models.User, error) {
	return s.r.GetEntity(ctx, id)
}

func (s userService) PostEntity(ctx context.Context, entity models.User) (models.User, error) {
	return s.r.PostEntity(ctx, entity)
}

func (s userService) GetEntities(ctx context.Context) ([]models.User, error) {
	return s.r.GetEntities(ctx)
}

func (s userService) PutEntity(ctx context.Context, id uuid.UUID, newEntity models.User) (bool, error) {
	return s.r.PutEntity(ctx, id, newEntity)
}

func (s userService) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.r.DeleteEntity(ctx, id)
}

func (s userService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return s.r.GetUserByEmail(ctx, email)
}
