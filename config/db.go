package config

import (
	"fmt"
	"os"

	"github.com/elpahlevi/go-restapi-boilerplate/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func dsnConfig() (string, error) {
	if os.Getenv("DB_HOST") == "" {
		return "", fmt.Errorf("DB_HOST env must be defined")
	}
	if os.Getenv("DB_PORT") == "" {
		return "", fmt.Errorf("DB_PORT env must be defined")
	}
	if os.Getenv("DB_USER") == "" {
		return "", fmt.Errorf("DB_USER env must be defined")
	}
	if os.Getenv("DB_DATABASE") == "" {
		return "", fmt.Errorf("DB_DATABASE env must be defined")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		return "", fmt.Errorf("DB_PASSWORD env must be defined")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"))

	return dsn, nil
}

func ConnectPg() error {
	dsn, err := dsnConfig()
	if err != nil {
		return err
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&model.Students{})
	return nil
}

func DBManager() *gorm.DB {
	return db
}
