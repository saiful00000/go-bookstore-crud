package main

import (
	"go-bootcamp/pkg/containers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	containers.Surve(e)
}