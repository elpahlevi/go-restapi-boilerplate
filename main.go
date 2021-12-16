package main

import (
	"fmt"
	"os"

	"github.com/elpahlevi/go-restapi-boilerplate/config"
	"github.com/elpahlevi/go-restapi-boilerplate/routes"
)

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		panic("SERVER_PORT env must be defined")
	}

	err := config.ConnectPg()
	if err != nil {
		panic("Cannot connect to database. \nReason: " + err.Error())
	} else {
		fmt.Println("Database Connected")
	}

	e := routes.Init()
	e.Logger.Fatal(e.Start(":" + serverPort))
}
