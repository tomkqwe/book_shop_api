package pgrepo

import (
	"book_shop_api/internal/domain"
	"book_shop_api/internal/repository/pgrepo/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepoInterface interface {
	CreateBook(model models.Book) (models.Book, error)
	GetBook(id uuid.UUID) (models.Book, error)
	GetBooks() ([]models.Book, error)
	UpdateBook(id uuid.UUID, newModel models.Book) (models.Book, error)
	DeleteBook(id uuid.UUID) (bool, error)
}

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepoInterface {
	return bookRepo{
		db: db,
	}
}

func (br bookRepo) CreateBook(modelBook models.Book) (models.Book, error) {
	modelBook.BeforeCreate(br.db)
	create := br.db.Create(&modelBook)
	if create.RowsAffected == 0 {
		return models.Book{}, domain.ErrCreateBook
	}

	return modelBook, nil

}

func (br bookRepo) GetBook(id uuid.UUID) (models.Book, error) {
	var model models.Book
	findByID := br.db.First(&model, id)
	if findByID.RowsAffected == 0 {
		return models.Book{}, domain.ErrFindBookById
	}

	return model, nil
}

func (br bookRepo) GetBooks() ([]models.Book, error) {
	var models []models.Book
	findAll := br.db.Find(&models)
	if findAll.RowsAffected == 0 {
		return nil, domain.ErrFindBooks
	}

	return models, nil
}

func (br bookRepo) UpdateBook(id uuid.UUID, newModel models.Book) (models.Book, error) {
	var oldModel models.Book
	br.db.First(&oldModel, id)
	mergeEntitys := mergeEntyties(oldModel, newModel)
	update := br.db.Updates(&mergeEntitys)
	if update.RowsAffected == 0 {
		return models.Book{}, domain.ErrUpdateBook
	}

	return br.GetBook(id)
}

func (br bookRepo) DeleteBook(id uuid.UUID) (bool, error) {
	delete := br.db.Delete(&models.Book{}, id)
	if delete.RowsAffected == 0 {
		return false, domain.ErrDeleteBook
	}
	return true, nil
}

func mergeEntyties(old models.Book, new models.Book) models.Book {
	return models.Book{
		Title:         new.Title,
		Author:        new.Author,
		YearPublished: new.YearPublished,
		Price:         new.Price,
		Category:      new.Category,
		Id:            old.Id,
		CreatedAt:     old.CreatedAt,
	}
}
