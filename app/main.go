package main

import (
	"fmt"
	"kredit-plus/domain"

	customerHandler "kredit-plus/customer/handler"
	customerRepo "kredit-plus/customer/repository"
	customerUsecase "kredit-plus/customer/usecase"

	customerTenorHandler "kredit-plus/customer_tenor/handler"
	customerTenorRepo "kredit-plus/customer_tenor/repository"
	customerTenorUsecase "kredit-plus/customer_tenor/usecase"

	transactionHandler "kredit-plus/transaction/handler"
	transactionRepo "kredit-plus/transaction/repository"
	transactionUsecase "kredit-plus/transaction/usecase"

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

	// Setup customerHandler.
	cRepo := customerRepo.NewCustomerRepository(DB)
	cUsecase := customerUsecase.NewCustomerUsecase(cRepo)
	customerHandler.NewCustomerHandler(server, cUsecase)

	// Setup customerTenorHandler.
	ctRepo := customerTenorRepo.NewCustomerTenorRepository(DB)
	ctUsecase := customerTenorUsecase.NewCustomerTenorUsecase(ctRepo, cRepo)
	customerTenorHandler.NewCustomerTenorHandler(server, ctUsecase)

	// Setup transactionHandler.
	tRepo := transactionRepo.NewTransactionRepository(DB)
	tUsecase := transactionUsecase.NewTransactionUsecase(tRepo)
	transactionHandler.NewTransactionHandler(server, tUsecase)

	server.Run(":8080")
}
