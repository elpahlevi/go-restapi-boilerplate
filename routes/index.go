package routes

import (
	"github.com/elpahlevi/go-restapi-boilerplate/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo) {
	// Index Routes
	e.GET("/", controller.GetIndex)

	// Routes for API
	api := e.Group("/api", middleware.Logger())
	// Routes for API endpoint /student. You can add any routes as you want
	StudentRoutes(api)
}
