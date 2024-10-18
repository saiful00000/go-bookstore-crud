package domain

import (
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
)

type IBookRepo interface {
	GetBooks(bookId uint) []models.Book
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookId uint) error
}

type IBookService interface {
	GetBooks(bookId uint) ([]types.BookRequest, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookId uint) error
}