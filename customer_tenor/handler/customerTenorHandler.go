package handler

import (
	"fmt"
	"kredit-plus/domain"
	"net/http"
	"strconv"

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
		routes.GET("/", handler.FetchAll)
		routes.GET("/:id", handler.FetchByID)
		routes.POST("", handler.Create)
		routes.PUT("/:id", handler.Update)
		routes.DELETE("/:id", handler.Delete)
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

func (handler *CustomerTenorHandler) Update(ginContext *gin.Context) {
	var reqBody domain.CustomerTenorUpdateRequest

	// Retrieve CustomerTenor ID from URL.
	uint64ID, err := strconv.ParseUint(ginContext.Param("id"), 10, 32)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid ID format"})
		return
	}

	var customerTenorID uint = uint(uint64ID)

	// Bind to struct.
	err = ginContext.ShouldBindJSON(&reqBody)
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
	customerTenorData, err := handler.Usecase.Update(ctx, &reqBody, customerTenorID)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// Return success response.
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   customerTenorData,
	})
}

func (handler *CustomerTenorHandler) FetchByID(ginContext *gin.Context) {
	// Retrieve CustomerTenor ID from URL.
	uint64ID, err := strconv.ParseUint(ginContext.Param("id"), 10, 32)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid ID format"})
		return
	}
	var customerTenorID uint = uint(uint64ID)

	// Pass to Usecase.
	ctx := ginContext.Request.Context()
	customerTenor, err := handler.Usecase.FetchByID(ctx, customerTenorID)
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

func (handler *CustomerTenorHandler) FetchAll(ginContext *gin.Context) {
	// Pass to Usecase.
	ctx := ginContext.Request.Context()
	customerTenors, err := handler.Usecase.FetchAll(ctx)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// Return success response.
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   customerTenors,
	})
}

func (handler *CustomerTenorHandler) Delete(ginContext *gin.Context) {
	// Retrieve CustomerTenor ID from URL.
	uint64ID, err := strconv.ParseUint(ginContext.Param("id"), 10, 32)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid ID format"})
		return
	}

	var customerTenorID uint = uint(uint64ID)

	// Pass to Usecase.
	ctx := ginContext.Request.Context()
	err = handler.Usecase.Delete(ctx, customerTenorID)

	// Return success response.
	ginContext.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Customer Tenor successfully deleted.",
	})
}
