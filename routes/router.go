package routes

import (
	"github.com/elpahlevi/go-restapi-boilerplate/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/api/students", controller.GetStudents)
	return e
}
