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

var customerTenorUsecase = &mocks.CustomerTenorUsecase{Mock: mock.Mock{}}
var customerTenorHandler = CustomerTenorHandler{
	Usecase: customerTenorUsecase,
}

var inputCreate = `{
    "CustomerID": 1,
    "Bulan": 11,
    "Limit": 200000
}`

var inputUpdate = `{
    "CustomerID": 1,
    "Bulan": 2,
    "Limit": 77700
}`

func TestCreateSuccess(t *testing.T) {
	body := []byte(inputCreate)
	req := httptest.NewRequest(http.MethodPost, "/api/customer/tenor", bytes.NewBuffer(body))
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	customerTenorUsecase.Mock.On("Create", mock.Anything, mock.Anything).Return(&domain.CustomerTenor{}, nil)

	customerTenorHandler.Create(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestUpdateSuccess(t *testing.T) {
	body := []byte(inputUpdate)
	req := httptest.NewRequest(http.MethodPut, "/api/customer/tenor/:id", bytes.NewBuffer(body))
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	customerTenorUsecase.Mock.On("Update", mock.Anything, mock.Anything).Return(&domain.CustomerTenor{}, nil)

	customerTenorHandler.Update(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestFetchByIDSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/customer/tenor/:id", nil)
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	customerTenorUsecase.Mock.On("FetchByID", mock.Anything, mock.Anything).Return(domain.CustomerTenor{}, nil)

	customerTenorHandler.FetchByID(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestFetchAllSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/customer/tenor", nil)
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	customerTenorUsecase.Mock.On("FetchAll", mock.Anything).Return([]domain.CustomerTenor{}, nil)

	customerTenorHandler.FetchAll(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestDeleteSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/customer/tenor/:id", nil)
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	customerTenorUsecase.Mock.On("Delete", mock.Anything).Return([]domain.CustomerTenor{}, nil)

	customerTenorHandler.Delete(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}
