package routes

import (
	"github.com/elpahlevi/go-restapi-boilerplate/controller"
	"github.com/labstack/echo/v4"
)

func StudentRoutes(eg *echo.Group) {
	eg.GET("/student", controller.GetStudents)
	eg.POST("/student", controller.InsertStudent)
}
