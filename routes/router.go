package routes

import (
	"github.com/elpahlevi/go-restapi-boilerplate/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", controller.GetIndex)
	e.GET("/api/students", controller.GetStudents)
	e.POST("/api/students", controller.InsertStudent)

	return e
}
