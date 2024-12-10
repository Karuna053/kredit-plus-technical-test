package main

import (
	"fmt"
	customerRoutes "kredit-plus/customer/route"
	"kredit-plus/domain"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Start database connection.
	var DB *gorm.DB
	dbHost := viper.GetString(`database.host`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.password`)
	dbName := viper.GetString(`database.name`)
	dbPort := viper.GetString(`database.port`)

	dbConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dbConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.Debug().AutoMigrate(
		&domain.Customer{},
		&domain.CustomerTenor{},
		&domain.Transaction{},
	)

	// Run Gin.
	server := gin.Default()
	customerRoutes.SetRoutes(server)

	server.Run(":8080")
}
