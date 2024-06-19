package domain

import "errors"

var (
	ErrRequired      = errors.New("required field")
	ErrNotFound      = errors.New("not found")
	ErrCreateBook    = errors.New("error create book")
	ErrFindBookById  = errors.New("error find book by id")
	ErrFindBooks     = errors.New("error find books")
	ErrUpdateBook    = errors.New("error update book")
	ErrDeleteBook    = errors.New("error delete book")
	ErrNil           = errors.New("data is nil")
	ErrIncorrectYear = errors.New("incorrect year")
	ErrPrice         = errors.New("incorrect price")
	ErrInvalidBookId = errors.New("invalid book ID")
)
