package handler

import (
	"fmt"
	"kredit-plus/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerTenorHandler struct {
	Usecase domain.CustomerTenorUsecase
}

func NewCustomerTenorHandler(ginEngine *gin.Engine, us domain.CustomerTenorUsecase) {
	handler := &CustomerTenorHandler{
		Usecase: us,
	}

	routes := ginEngine.Group("/api/customer/tenor")
	{
		routes.POST("", handler.Create)
	}
}

func (handler *CustomerTenorHandler) Create(ginContext *gin.Context) {
	var reqBody domain.CustomerTenorCreateRequest

	// Bind to struct.
	err := ginContext.ShouldBindJSON(&reqBody)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// Validate struct.
	validate := validator.New()
	err = validate.Struct(reqBody)

	if err != nil {
		// Extracting validation errors
		errorDetails := make(map[string]string)
		for _, validationErr := range err.(validator.ValidationErrors) {
			errorDetails[validationErr.Field()] = fmt.Sprintf("Validation failed on '%s' tag", validationErr.Tag())
		}

		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errorDetails})
		return
	}

	// Pass to Usecase.
	ctx := ginContext.Request.Context()
	customerTenor, err := handler.Usecase.Create(ctx, &reqBody)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// Return success response.
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   customerTenor,
	})
}
