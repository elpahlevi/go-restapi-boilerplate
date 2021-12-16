package controller

import (
	"net/http"

	"github.com/elpahlevi/go-restapi-boilerplate/config"
	"github.com/elpahlevi/go-restapi-boilerplate/model"

	"github.com/labstack/echo/v4"
)

func GetIndex(c echo.Context) error {
	welcomeText := "Welcome to go-restapi-boilerplate. \nTo continue, access the endpoint /api/students"
	return c.String(http.StatusOK, welcomeText)
}

func GetStudents(c echo.Context) error {
	db := config.DBManager()
	students := new([]model.Students)

	db.Find(&students)
	return c.JSON(http.StatusOK, students)
}

func InsertStudent(c echo.Context) error {
	db := config.DBManager()

	student := new(model.Students)
	if err := c.Bind(student); err != nil {
		return err
	}
	if student.Name == "" {
		return c.String(http.StatusInternalServerError, "Please enter your name")
	}

	db.Create(&student)
	return c.JSON(http.StatusOK, "One student added")
}
