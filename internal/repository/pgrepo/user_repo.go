package pgrepo

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"book_shop_api/internal/transport/utils"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ CRUDRepository[models.User] = (*userRepo)(nil)

type UserRepository interface {
	CRUDRepository[models.User] // Включаем методы из параметризированного репозитория
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) UserRepository {
	return userRepo{db: db}
}

func (r userRepo) PostEntity(ctx context.Context, model models.User) (models.User, error) {
	sql := `INSERT INTO users(email,password,role) VALUES ($1,$2,'user') RETURNING id`
	var id uuid.UUID
	hashPass, err := utils.HashPassword(model.Password)
	if err != nil {
		return models.User{}, err
	}

	err = r.db.QueryRow(ctx, sql, model.Email, hashPass).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	model.ID = id

	return model, nil
}

func (r userRepo) GetEntity(ctx context.Context, id uuid.UUID) (models.User, error) {
	sql := `SELECT id,email,password,role FROM users WHERE id = $1 LIMIT 1`
	var user models.User
	err := r.db.QueryRow(ctx, sql, id).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r userRepo) GetEntities(ctx context.Context) ([]models.User, error) {
	sql := `SELECT id,email,password,role FROM users`
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r userRepo) PutEntity(ctx context.Context, id uuid.UUID, newModel models.User) (bool, error) {
	sql := `UPDATE users SET email = $1,password=$2,role= $3 WHERE id = $4`
	res, err := r.db.Exec(ctx, sql, newModel.Email, newModel.Password, newModel.Role, id)
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}

func (r userRepo) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	sql := `DELETE FROM users WHERE id = $1`
	res, err := r.db.Exec(ctx, sql, id)
	if err != nil {
		return false, nil
	}

	return res.RowsAffected() > 0, nil
}

func (r userRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	sql := `SELECT id,email,password,role FROM users WHERE email = $1 LIMIT 1`
	var user models.User
	err := r.db.QueryRow(ctx, sql, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
