package pgrepo

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ CRUDRepository[models.Order] = (*orderRepo)(nil)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) orderRepo {
	return orderRepo{db: db}
}

func (r orderRepo) PostEntity(ctx context.Context, model models.Order) (models.Order, error) {
	sql := `INSERT INTO orders (user_id,order_date,total_amount) VALUES ($1,$2,$3) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, sql, model.UserID, time.Now(), model.TotalAmount).Scan(&id)
	if err != nil {
		return models.Order{}, err
	}

	model.ID = id

	return model, nil
}

func (r orderRepo) GetEntity(ctx context.Context, id uuid.UUID) (models.Order, error) {
	sql := `SELECT id,user_id,order_date,total_amount FROM orders WHERE id = $1 LIMIT 1`
	var order models.Order
	err := r.db.QueryRow(ctx, sql, id).Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (r orderRepo) GetEntities(ctx context.Context) ([]models.Order, error) {
	sql := `SELECT id,user_id,order_date,total_amount FROM orders`
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := make([]models.Order, 0)
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r orderRepo) PutEntity(ctx context.Context, id uuid.UUID, newModel models.Order) (bool, error) {
	sql := `UPDATE orders SET user_id=$1,order_date=$2,total_amount=$3 WHERE id = $4`
	res, err := r.db.Exec(ctx, sql, newModel.UserID, time.Now(), newModel.TotalAmount, id)
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}

func (r orderRepo) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	sql := `DELETE FROM orders WHERE id = $1`
	res, err := r.db.Exec(ctx, sql, id)
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}
