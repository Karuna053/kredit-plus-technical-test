package controller

import (
	"fmt"
	"kredit-plus/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	// gin *gin.Engine
}

func NewCustomerHandler(g *gin.Engine) {
	handler := &CustomerHandler{}

	routes := g.Group("/api/customer")
	{
		// Customer
		routes.POST("/create", handler.CreateCustomer)
	}
}

func (h CustomerHandler) CreateCustomer(c *gin.Context) {
	var DB *gorm.DB
	var customer domain.Customer

	// Parse JSON request and populate Customer struct
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
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
		NIK:          req.NIK,
		FullName:     req.FullName,
		LegalName:    req.LegalName,
		TempatLahir:  req.TempatLahir,
		TanggalLahir: req.TanggalLahir,
		Gaji:         req.Gaji,
		FotoKTP:      req.FotoKTP,
		FotoSelfie:   req.FotoSelfie,
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
