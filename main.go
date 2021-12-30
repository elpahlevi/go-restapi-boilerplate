package main

import (
	"fmt"
	"os"

	"github.com/elpahlevi/go-restapi-boilerplate/config"
	"github.com/elpahlevi/go-restapi-boilerplate/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		panic("SERVER_PORT env must be defined")
	}

	// Initialize Echo instance
	e := echo.New()

	// Connect to database
	err := config.ConnectPg()
	if err != nil {
		panic("Cannot connect to database. \nReason: " + err.Error())
	} else {
		fmt.Println("Database Connected")
	}

	// Load Routes
	routes.Setup(e)
	// Start the server
	e.Logger.Fatal(e.Start(":" + serverPort))
}
