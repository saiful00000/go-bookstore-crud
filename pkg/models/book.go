package models

import "go-bootcamp/pkg/types"

type Book struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	BookName    string
	Author      string
	Publication string
}

func (book *Book) ToBookRequest() types.BookRequest {
	return types.BookRequest{
		ID: book.ID,
		BookName: book.BookName,
		Author: book.Author,
		Publication: book.Publication,
	}
} 