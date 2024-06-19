package pgrepo

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"testing"

	"github.com/go-kiss/monkey"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNewBookRepo(t *testing.T) {
	//given
	db := &gorm.DB{}
	expected := bookRepo{db: db}

	//when
	actual := NewBookRepo(db)

	//then
	require.Equal(t, expected, actual)
}

func TestCreateBookSuccess(t *testing.T) {
	//given
	br := bookRepo{}
	expected := models.Book{}
	monkey.Patch((*models.Book).BeforeCreate, func(*models.Book, *gorm.DB) error {
		return nil
	})
	monkey.Patch((*gorm.DB).Create, func(*gorm.DB, interface{}) *gorm.DB {
		return &gorm.DB{RowsAffected: 1}
	})

	//when
	actual, err := br.CreateBook(expected)
	monkey.UnpatchAll()

	//then
	require.Equal(t, expected, actual)
	require.NoError(t, err)
}

func TestCreateBookErr(t *testing.T) {
	//given
	br := bookRepo{}
	expected := models.Book{}
	monkey.Patch((*gorm.DB).Create, func(*gorm.DB, interface{}) *gorm.DB {
		return &gorm.DB{}
	})

	//when
	actual, err := br.CreateBook(expected)
	monkey.UnpatchAll()

	//then
	require.Equal(t, expected, actual)
	require.Error(t, err)
}

func TestGetBookSuccess(t *testing.T) {
	//given
	br := bookRepo{}
	expected := models.Book{}
	monkey.Patch((*gorm.DB).First, func(*gorm.DB, interface{}, ...interface{}) *gorm.DB {
		return &gorm.DB{RowsAffected: 1}
	})

	//when
	actual, err := br.GetBook(uuid.Nil)
	monkey.UnpatchAll()

	//then
	require.NoError(t, err)
	require.Equal(t, expected, actual)

}

func TestGetBookError(t *testing.T) {
	//given
	br := bookRepo{}
	expected := models.Book{}
	monkey.Patch((*gorm.DB).First, func(*gorm.DB, interface{}, ...interface{}) *gorm.DB {
		return &gorm.DB{}
	})

	//when
	actual, err := br.GetBook(uuid.Nil)
	monkey.UnpatchAll()

	//then
	require.Error(t, err)
	require.Equal(t, expected, actual)

}

func TestGetBooksSuccess(t *testing.T) {
	//given
	br := bookRepo{}
	monkey.Patch((*gorm.DB).Find, func(*gorm.DB, interface{}, ...interface{}) *gorm.DB {
		return &gorm.DB{RowsAffected: 1}
	})

	//when
	actual, err := br.GetBooks()
	monkey.UnpatchAll()
	//then
	require.NoError(t, err)
	require.Nil(t, actual)
}

func TestGetBooksError(t *testing.T) {
	//given
	br := bookRepo{}
	monkey.Patch((*gorm.DB).Find, func(*gorm.DB, interface{}, ...interface{}) *gorm.DB {
		return &gorm.DB{}
	})

	//when
	actual, err := br.GetBooks()
	monkey.UnpatchAll()

	//then
	require.Error(t, err)
	require.Nil(t, actual)
}

func TestUpdateBookSuccess(t *testing.T) {
	//given
	br := bookRepo{}
	expected := models.Book{
		Title:         "a",
		Author:        "b",
		YearPublished: 1,
		Price:         decimal.NewFromInt(22),
		Category:      "c",
	}
	monkey.Patch((*gorm.DB).First, func(*gorm.DB, interface{}, ...interface{}) *gorm.DB {
		return &gorm.DB{RowsAffected: 1}
	})
	monkey.Patch((*gorm.DB).Updates, func(*gorm.DB, interface{}) *gorm.DB {
		return &gorm.DB{RowsAffected: 1}
	})
	monkey.Patch(bookRepo.GetBook, func(bookRepo, uuid.UUID) (models.Book, error) {
		return expected, nil
	})

	//when
	actual, err := br.UpdateBook(uuid.Nil, expected)
	monkey.UnpatchAll()

	//then
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestUpdateBookErr(t *testing.T) {
	//given
	br := bookRepo{}
	monkey.Patch((*gorm.DB).First, func(*gorm.DB, interface{}, ...interface{}) *gorm.DB {
		return &gorm.DB{}
	})
	monkey.Patch((*gorm.DB).Updates, func(*gorm.DB, interface{}) *gorm.DB {
		return &gorm.DB{}
	})

	//when
	actual, err := br.UpdateBook(uuid.Nil, models.Book{})
	monkey.UnpatchAll()
	//then
	require.Error(t, err)
	require.Equal(t, models.Book{}, actual)
}

func TestDeleteBookSuccess(t *testing.T) {
	//given
	br := bookRepo{}
	monkey.Patch((*gorm.DB).Delete, func(*gorm.DB, interface{}, ...interface{}) *gorm.DB {
		return &gorm.DB{RowsAffected: 1}
	})

	//when
	actual, err := br.DeleteBook(uuid.Nil)
	monkey.UnpatchAll()

	//then
	require.Equal(t, true, actual)
	require.NoError(t, err)
}

func TestDeleteBookErr(t *testing.T) {
	//given
	br := bookRepo{}
	monkey.Patch((*gorm.DB).Delete, func(*gorm.DB, interface{}, ...interface{}) *gorm.DB {
		return &gorm.DB{}
	})

	//when
	actual, err := br.DeleteBook(uuid.Nil)
	monkey.UnpatchAll()

	//then
	require.Equal(t, false, actual)
	require.Error(t, err)
}

func TestMergeEntyties(t *testing.T) {
	//given
	expected := models.Book{
		Title:         "a",
		Author:        "b",
		YearPublished: 1,
		Price:         decimal.NewFromInt(22),
		Category:      "c",
	}
	mb := models.Book{}

	//when
	actual := mergeEntyties(mb, expected)

	//then
	require.Equal(t, expected, actual)
}
