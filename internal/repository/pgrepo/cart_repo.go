package pgrepo

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ CRUDRepository[models.Cart] = (*cartRepo)(nil)

type cartRepo struct {
	db *pgxpool.Pool
}

func NewCartRepo(db *pgxpool.Pool) cartRepo {
	return cartRepo{db: db}
}

func (r cartRepo) PostEntity(ctx context.Context, model models.Cart) (models.Cart, error) {
	sql := `INSERT INTO carts (user_id,created_at) VALUES ($1,NOW()) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, sql, model.UserID).Scan(&id)
	if err != nil {
		return models.Cart{}, err
	}
	model.ID = id

	return model, nil
}

func (r cartRepo) GetEntity(ctx context.Context, id uuid.UUID) (models.Cart, error) {
	sql := `SELECT id,user_id,created_at FROM carts WHERE id = $1 LIMIT 1`
	var cart models.Cart
	err := r.db.QueryRow(ctx, sql, id).Scan(&cart.ID, &cart.UserID, &cart.CreatedAt)
	if err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func (r cartRepo) GetEntities(ctx context.Context) ([]models.Cart, error) {
	sql := `SELECT id,user_id,created_at FROM carts`
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	carts := make([]models.Cart, 0)
	for rows.Next() {
		var cart models.Cart
		if err := rows.Scan(&cart.ID, &cart.UserID, &cart.CreatedAt); err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}

	return carts, nil
}

func (r cartRepo) PutEntity(ctx context.Context, id uuid.UUID, newModel models.Cart) (bool, error) {
	sql := `UPDATE carts SET created_at = $1,user_id = $2 WHERE id = $3`
	result, err := r.db.Exec(ctx, sql, time.Now(), newModel.UserID, id)
	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}

func (r cartRepo) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	sql := `DELETE FROM carts WHERE id = $1`
	res, err := r.db.Exec(ctx, sql, id)
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}
