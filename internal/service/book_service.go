package service

import (
	"book_shop_api/internal/repository/pgrepo"
	"book_shop_api/internal/repository/pgrepo/models"

	"github.com/google/uuid"
)

type BookServiceInterface interface {
	GetBook(id uuid.UUID) (models.Book, error)
	CreateBook(mb models.Book) (models.Book, error)
	GetBooks() ([]models.Book, error)
	UpdateBook(id uuid.UUID, newModel models.Book) (models.Book, error)
	DeleteBook(id uuid.UUID) (bool, error)
}

type bookService struct {
	pgrepo pgrepo.BookRepoInterface
}

func NewBookService(pgrepos pgrepo.BookRepoInterface) BookServiceInterface {
	return bookService{pgrepo: pgrepos}
}

// CreateBook implements BookServiceInterface.
func (b bookService) CreateBook(mb models.Book) (models.Book, error) {
	return b.pgrepo.CreateBook(mb)
}

// DeleteBook implements BookServiceInterface.
func (b bookService) DeleteBook(id uuid.UUID) (bool, error) {
	return b.pgrepo.DeleteBook(id)
}

// GetBook implements BookServiceInterface.
func (b bookService) GetBook(id uuid.UUID) (models.Book, error) {
	return b.pgrepo.GetBook(id)

}

// GetBooks implements BookServiceInterface.
func (b bookService) GetBooks() ([]models.Book, error) {
	return b.pgrepo.GetBooks()
}

// UpdateBook implements BookServiceInterface.
func (b bookService) UpdateBook(id uuid.UUID, newModel models.Book) (models.Book, error) {
	return b.pgrepo.UpdateBook(id, newModel)
}
