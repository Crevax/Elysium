package controllers

import (
	"net/http"

	"cjdavis.me/elysium/services"
	"github.com/labstack/echo"
)

type ProfileController struct{}

func (ctrl *ProfileController) Init(router *echo.Echo) {
	router.GET("/", Index)
}

func Index(c echo.Context) error {
	p := services.GetProfileService().GetProfile()

	return c.JSON(http.StatusOK, p)
}
