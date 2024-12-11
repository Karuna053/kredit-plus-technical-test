package handler

import (
	"fmt"
	"kredit-plus/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransactionHandler struct {
	Usecase domain.TransactionUsecase
}

func NewTransactionHandler(ginEngine *gin.Engine, us domain.TransactionUsecase) {
	handler := &TransactionHandler{
		Usecase: us,
	}

	routes := ginEngine.Group("/api/transaction")
	{
		routes.GET("/", handler.FetchAll)
		routes.GET("/:id", handler.FetchByID)
		routes.POST("", handler.Create)
		routes.PUT("/:id", handler.Update)
	}
}

func (handler *TransactionHandler) Create(ginContext *gin.Context) {
	var reqBody domain.TransactionCreateRequest

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
	transactionData, err := handler.Usecase.Create(ctx, &reqBody)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// Return success response.
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   transactionData,
	})
}

func (handler *TransactionHandler) Update(ginContext *gin.Context) {
	var reqBody domain.TransactionUpdateRequest

	// Retrieve Transaction ID from URL.
	paramID, err := strconv.ParseUint(ginContext.Param("id"), 10, 32)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid ID format"})
		return
	}

	var transactionID uint = uint(paramID)

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
	transactionData, err := handler.Usecase.Update(ctx, &reqBody, transactionID)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// Return success response.
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   transactionData,
	})
}

func (handler *TransactionHandler) FetchByID(ginContext *gin.Context) {
	// Retrieve Transaction ID from URL.
	paramID, err := strconv.ParseUint(ginContext.Param("id"), 10, 32)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid ID format"})
		return
	}
	var transactionID uint = uint(paramID)

	// Pass to Usecase.
	ctx := ginContext.Request.Context()
	transaction, err := handler.Usecase.FetchByID(ctx, transactionID)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// Return success response.
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   transaction,
	})
}

func (handler *TransactionHandler) FetchAll(ginContext *gin.Context) {
	// Pass to Usecase.
	ctx := ginContext.Request.Context()
	transactions, err := handler.Usecase.FetchAll(ctx)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// Return success response.
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   transactions,
	})
}
