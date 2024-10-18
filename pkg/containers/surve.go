package containers

import (
	"fmt"
	"go-bootcamp/pkg/config"
	"go-bootcamp/pkg/connection"
	"go-bootcamp/pkg/controllers"
	"go-bootcamp/pkg/repositories"
	"go-bootcamp/pkg/routes"
	"go-bootcamp/pkg/services"
	"log"

	"github.com/labstack/echo/v4"
)

func Surve(e *echo.Echo) {
	config.SetCOnfig()

	db := connection.GetDb()

	// book related instances
	bookRepo := repositories.BookDbInstance(db)
	bookService := services.BookServiceInstance(bookRepo)
	controllers.SetBookService(bookService)

	routes.BookRoutes(e)
	// Start the server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}