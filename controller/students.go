package controller

import (
	"net/http"

	"github.com/elpahlevi/go-restapi-boilerplate/config"
	"github.com/elpahlevi/go-restapi-boilerplate/model"

	"github.com/labstack/echo/v4"
)

func GetStudents(c echo.Context) error {
	students := []model.Students{}
	db := config.DBManager()

	db.Find(&students) //select * from students;. The result will be append to variable students
	return c.JSON(http.StatusOK, students)
}
