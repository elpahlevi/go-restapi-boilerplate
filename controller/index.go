package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetIndex(c echo.Context) error {
	welcomeText := "Welcome to go-restapi-boilerplate. \nTo continue, access the endpoint /api/students"
	return c.String(http.StatusOK, welcomeText)
}
