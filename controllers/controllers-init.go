package controllers

import "github.com/labstack/echo"

func Init(router *echo.Echo) {
	new(ProfileController).Init(router)
	new(AuthorController).Init(router)
	new(BookController).Init(router)
}
