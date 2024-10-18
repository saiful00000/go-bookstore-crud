package services

import (
	"errors"
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
)

type BookService struct {
	repo domain.IBookRepo
}

func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &BookService{
		repo: bookRepo,
	}
}

func (service *BookService) CreateBook(book *models.Book) error {
	if err := service.repo.CreateBook(book); err != nil {
		return errors.New("Failed to creating the book")
	}

	return nil
}


func (service *BookService) GetBooks(bookId uint) ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	books := service.repo.GetBooks(bookId)
	if len(books) == 0 {
		return nil, errors.New("book does not exists")
	}
	for _, book := range(books) {
		allBooks = append(allBooks, book.ToBookRequest())
	}

	return allBooks, nil
}

func (service *BookService) DeleteBook(bookId uint) error {return nil} 
func (service *BookService) UpdateBook(book *models.Book) error {return nil}

