package controllers

import (
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var BookService domain.IBookService

func SetBookService (bookService domain.IBookService) {
	BookService = bookService
}


func CreateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}
	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}

	book := &models.Book{
		BookName: reqBook.BookName,
		Author: reqBook.Author,
		Publication: reqBook.Publication,

	}

	if err := BookService.CreateBook(book); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "Book was created successfully")
}
func GetBooks(e echo.Context) error {
	tempBookId := e.QueryParam("bookID")
	bookId, err := strconv.ParseInt(tempBookId, 0, 0)

	// If the temp book id is empty that means no book id has given by client 
	if tempBookId != "" {
		if err != nil {
			return e.JSON(http.StatusBadRequest, "Enter a valid book Id")
		}
	}

	book, err := BookService.GetBooks(uint(bookId))

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, book)
}
func UpdateBook(e echo.Context) error {return nil}
func DeleteBook(e echo.Context) error {return nil}
