package controllers

import (
	"net/http"

	"cjdavis.me/elysium/services"

	"github.com/labstack/echo"
)

type BookController struct{}

func (ctrl *BookController) Init(router *echo.Echo) {
	router.GET("/book", GetAllBooks)
}

func GetAllBooks(c echo.Context) error {
	books := services.GetBookService().GetAllBooks()

	return c.JSON(http.StatusOK, books)
}
