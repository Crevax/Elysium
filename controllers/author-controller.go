package controllers

import (
	"net/http"
	"strconv"

	"cjdavis.me/elysium/services"

	"github.com/labstack/echo"
)

type AuthorController struct{}

func (ctrl *AuthorController) Init(router *echo.Echo) {
	router.GET("/author", GetAllAuthors)
	router.GET("/author/:id/book", GetAllBooksForAuthor)
}

func GetAllAuthors(c echo.Context) error {
	authors := services.GetAuthorService().GetAllAuthors()

	return c.JSON(http.StatusOK, authors)
}

func GetAllBooksForAuthor(c echo.Context) error {
	authorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error)
	}

	books := services.GetBookService().GetBooksByAuthor(authorID)

	return c.JSON(http.StatusOK, books)
}
