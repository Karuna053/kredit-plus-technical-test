package handler

import (
	"bytes"
	"kredit-plus/domain"
	"kredit-plus/domain/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transactionUsecase = &mocks.TransactionUsecase{Mock: mock.Mock{}}
var transactionHandler = TransactionHandler{
	Usecase: transactionUsecase,
}

var inputCreate = `{
    "CustomerID": 1,
    "AdminFee": 2000,
    "JumlahBunga": 4000,
    "JumlahCicilan": 170000,
    "NomorKontrak": "Hello hello",
    "NomorOnTheRoad": "Hullo Hullo"
}`

var inputUpdate = `{
    "CustomerID": 1,
    "AdminFee": 3000,
    "JumlahBunga": 4000,
    "JumlahCicilan": 10000,
    "NomorKontrak": "Millian",
    "NomorOnTheRoad": "Terrible"
}`

func TestCreateSuccess(t *testing.T) {
	body := []byte(inputCreate)
	req := httptest.NewRequest(http.MethodPost, "/api/transaction", bytes.NewBuffer(body))
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	transactionUsecase.Mock.On("Create", mock.Anything, mock.Anything).Return(&domain.Transaction{}, nil)

	transactionHandler.Create(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestUpdateSuccess(t *testing.T) {
	body := []byte(inputUpdate)
	req := httptest.NewRequest(http.MethodPut, "/api/transaction/:id", bytes.NewBuffer(body))
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	transactionUsecase.Mock.On("Update", mock.Anything, mock.Anything).Return(&domain.Transaction{}, nil)

	transactionHandler.Update(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestFetchByIDSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/transaction/:id", nil)
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	transactionUsecase.Mock.On("FetchByID", mock.Anything, mock.Anything).Return(domain.Transaction{}, nil)

	transactionHandler.FetchByID(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestFetchAllSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/transaction", nil)
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	transactionUsecase.Mock.On("FetchAll", mock.Anything).Return([]domain.Transaction{}, nil)

	transactionHandler.FetchAll(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}
