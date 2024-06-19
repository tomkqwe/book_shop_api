package pgrepo

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ CRUDRepository[models.OrderItem] = (*orderItems)(nil)

type orderItems struct {
	db *pgxpool.Pool
}

func NewOrderItemRepo(db *pgxpool.Pool) orderItems {
	return orderItems{db: db}
}

func (r orderItems) PostEntity(ctx context.Context, model models.OrderItem) (models.OrderItem, error) {
	sql := `INSERT INTO order_items(order_id ,book_id ,quantity ,price) VALUES ($1,$2,$3,$4) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, sql, model.OrderID, model.BookID, model.Quantity, model.Price).Scan(&id)
	if err != nil {
		return models.OrderItem{}, err
	}
	model.ID = id

	return model, nil
}

func (r orderItems) GetEntity(ctx context.Context, id uuid.UUID) (models.OrderItem, error) {
	sql := `SELECT id,order_id ,book_id ,quantity ,price FROM order_items WHERE id = $1 LIMIT 1`
	var orderItem models.OrderItem
	err := r.db.QueryRow(ctx, sql, id).Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.BookID, &orderItem.Quantity, &orderItem.Price)
	if err != nil {
		return models.OrderItem{}, err
	}

	return orderItem, nil
}

func (r orderItems) GetEntities(ctx context.Context) ([]models.OrderItem, error) {
	sql := `SELECT id,order_id ,book_id ,quantity ,price FROM order_items`
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]models.OrderItem, 0)
	for rows.Next() {
		var orderItem models.OrderItem
		if err := rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.BookID, &orderItem.Quantity, &orderItem.Price); err != nil {
			return nil, err
		}
		users = append(users, orderItem)
	}

	return users, nil
}

func (r orderItems) PutEntity(ctx context.Context, id uuid.UUID, newModel models.OrderItem) (bool, error) {
	sql := `UPDATE order_items SET order_id = $1 ,book_id=$2 ,quantity=$3 ,price=$4 WHERE id = $5`
	res, err := r.db.Exec(ctx, sql, newModel.OrderID, newModel.BookID, newModel.Quantity, newModel.Price, id)
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}

func (r orderItems) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	sql := `DELETE FROM order_items WHERE id = $1`
	res, err := r.db.Exec(ctx, sql, id)
	log.Printf("%v\n", err)
	if err != nil {
		return false, nil
	}

	return res.RowsAffected() > 0, nil
}
