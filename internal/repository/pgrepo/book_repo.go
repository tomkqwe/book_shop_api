package pgrepo

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ CRUDRepository[models.Book] = (*bookRepo)(nil)

type bookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) bookRepo {
	return bookRepo{
		db: db,
	}
}

func (r bookRepo) PostEntity(ctx context.Context, model models.Book) (models.Book, error) {
	sql := `INSERT INTO books (title,author,year_published,price,category) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, sql, model.Title, model.Author, model.YearPublished, model.Price, model.Category).Scan(&id)
	if err != nil {
		return models.Book{}, err
	}
	model.Id = id

	return model, nil
}

func (r bookRepo) GetEntity(ctx context.Context, id uuid.UUID) (models.Book, error) {
	sql := `SELECT id,title,author,year_published,price,category FROM books WHERE id = $1 LIMIT 1`
	var book models.Book
	err := r.db.QueryRow(ctx, sql, id).Scan(&book.Id, &book.Title, &book.Author, &book.YearPublished, &book.Price, &book.Category)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r bookRepo) GetEntities(ctx context.Context) ([]models.Book, error) {
	sql := `SELECT id,title,author,year_published,price,category FROM books`
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := make([]models.Book, 0)
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.YearPublished, &book.Price, &book.Category); err != nil {
			return nil, err
		}
		books = append(books, book)

	}

	return books, nil
}

func (r bookRepo) PutEntity(ctx context.Context, id uuid.UUID, newModel models.Book) (bool, error) {
	sql := `UPDATE books SET title = $1,author=$2,year_published=$3,price=$4,category=$5 WHERE id = $6`
	result, err := r.db.Exec(ctx, sql, newModel.Title, newModel.Author, newModel.YearPublished, newModel.Price, newModel.Category, id)
	if err != nil {
		return false, err
	}
	return result.RowsAffected() > 0, nil
}

func (r bookRepo) DeleteEntity(ctx context.Context, id uuid.UUID) (bool, error) {
	sql := `DELETE FROM books WHERE id = $1`
	res, err := r.db.Exec(ctx, sql, id)
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}
