package handler

import (
	"fmt"
	"kredit-plus/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerHandler struct {
	Usecase domain.CustomerUsecase
}

func NewCustomerHandler(ginEngine *gin.Engine) {
	handler := &CustomerHandler{}

	routes := ginEngine.Group("/api/customer")
	{
		routes.POST("", handler.Create)
		// routes.PUT("/:id", handler.Update)
	}
}

func (handler *CustomerHandler) Create(ginContext *gin.Context) {
	var reqBody domain.CustomerCreateRequest

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
	customerData, err := handler.Usecase.Create(ctx, &reqBody)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// Return success response.
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   customerData,
	})
}
