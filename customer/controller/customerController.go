package controller

import (
	"fmt"
	"kredit-plus/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CreateCustomer(c *gin.Context) {
	var DB *gorm.DB
	var customer domain.Customer

	// Parse JSON request and populate Customer struct
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validare request on create customer context
	var CustomerCreateRules domain.CustomerCreateRules
	CustomerCreateRules.NIK = customer.NIK
	CustomerCreateRules.FullName = customer.FullName
	CustomerCreateRules.LegalName = customer.LegalName
	CustomerCreateRules.TempatLahir = customer.TempatLahir
	CustomerCreateRules.TanggalLahir = customer.TanggalLahir
	CustomerCreateRules.Gaji = customer.Gaji
	CustomerCreateRules.FotoKTP = customer.FotoKTP
	CustomerCreateRules.FotoSelfie = customer.FotoSelfie

	validate := validator.New()
	err = validate.Struct(CustomerCreateRules)
	fmt.Println(err) // Logging error on console... just because.

	if err != nil {
		// Extracting validation errors
		errorDetails := make(map[string]string)
		for _, validationErr := range err.(validator.ValidationErrors) {
			errorDetails[validationErr.Field()] = fmt.Sprintf("Validation failed on '%s' tag", validationErr.Tag())
		}

		// Return error.
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": errorDetails,
		})
		return
	}

	// Create customer.
	customerInput := domain.Customer{
		NIK:          customer.NIK,
		FullName:     customer.FullName,
		LegalName:    customer.LegalName,
		TempatLahir:  customer.TempatLahir,
		TanggalLahir: customer.TanggalLahir,
		Gaji:         customer.Gaji,
		FotoKTP:      customer.FotoKTP,
		FotoSelfie:   customer.FotoSelfie,
	}

	err = DB.Create(&customerInput).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err,
		})
		return
	}

	// Return success response.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   customerInput,
	})
}
