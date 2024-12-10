package main

import (
	"fmt"
	"kredit-plus/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBHost = "localhost"
	DBUser = "postgres"
	DBPass = "password"
	DBName = "kredit-plus"
	DBPort = "5432"

	DB *gorm.DB // Declare data type.
)

func init() {
	// Initialize Database.
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost, DBUser, DBPass, DBName, DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.Debug().AutoMigrate(
		&domain.Customer{},
		&domain.CustomerTenor{},
		&domain.Transaction{},
	)
}
