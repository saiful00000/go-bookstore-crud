package repositories

import (
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"

	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func BookDbInstance(d *gorm.DB) domain.IBookRepo {
	return &bookRepo{
		db: d,
	}
}


func (repo *bookRepo) CreateBook(book *models.Book) error {
	if err := repo.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *bookRepo) GetBooks(bookId uint) []models.Book {
	var books []models.Book
	var err error

	if bookId != 0 {
		err = repo.db.Where("id = ?", bookId).Find(&books).Error
	} else {
		err = repo.db.Find(&books).Error
	}

	if err != nil {
		return []models.Book{}
	}

	return books
}

func (repo *bookRepo) DeleteBook(bookId uint) error {return nil}
func (repo *bookRepo) UpdateBook(book *models.Book) error {return nil}
