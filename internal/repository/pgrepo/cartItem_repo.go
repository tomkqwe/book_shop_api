package pgrepo

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ CRUDRepository[models.CartItem] = (*cartItemRepo)(nil)

type cartItemRepo struct {
	db *pgxpool.Pool
}

func NewCartItemRepo(db *pgxpool.Pool) cartItemRepo {
	return cartItemRepo{db: db}
}

func (r cartItemRepo) PostEntity(ctx context.Context, model models.CartItem) (models.CartItem, error) {
	sql := `INSERT INTO cart_items (cart_id,book_id,count) VALUES ($1,$2,$3) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, sql, model.CartID, model.BookID, model.Count).Scan(&id)
	if err != nil {
		return models.CartItem{}, err
	}

	model.ID = id

	return model, nil
}

func (r cartItemRepo) GetEntity(ctx context.Context, id uuid.UUID) (models.CartItem, error) {
	sql := `SELECT id,cart_id,book_id,count FROM cart_items WHERE id = $1 LIMIT 1`
	var cartItem models.CartItem
	err := r.db.QueryRow(ctx, sql, id).Scan(&cartItem.ID, &cartItem.CartID, &cartItem.BookID, &cartItem.Count)
	if err != nil {
		return models.CartItem{}, err
	}

	return cartItem, nil
}

func (r cartItemRepo) GetEntities(ctx context.Context) ([]models.CartItem, error) {
	sql := `SELECT id,cart_id,book_id,count From cart_items`
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cartItems := make([]models.CartItem, 0)
	for rows.Next() {
		var cartItem models.CartItem
		if err := rows.Scan(&cartItem.ID, &cartItem.CartID, &cartItem.BookID, &cartItem.Count); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func (r cartItemRepo) PutEntity(ctx context.Context, id uuid.UUID, newModel models.CartItem) (bool, error) {
	sql := `UPDATE cart_items SET cart_id=$1,book_id=$2,count=$3 WHERE id =$4`
	res, err := r.db.Exec(ctx, sql, newModel.CartID, newModel.BookID, newModel.Count, id)
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}

func (r cartItemRepo) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	sql := `DELETE FROM cart_items WHERE id = $1`
	res, err := r.db.Exec(ctx, sql, id)
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}
