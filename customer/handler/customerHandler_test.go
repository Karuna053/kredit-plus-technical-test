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

var customerUsecase = &mocks.CustomerUsecase{Mock: mock.Mock{}}
var customerHandler = CustomerHandler{
	Usecase: customerUsecase,
}

var inputCreate = `{
    "NIK": "123002122",
    "FotoKTP": "whatsup",
    "FotoSelfie": "whatsup 222",
    "FullName": "Eilang Eilang",
    "Gaji": 1000000,
    "LegalName": "Kario",
    "TanggalLahir": "2024-10-18",
    "TempatLahir": "Let's gooo"
}`

var inputUpdate = `{
    "NIK": "12322",
    "FotoKTP": "whatsupthree",
    "FotoSelfie": "whatsup 100",
    "FullName": "Him Hem Ham",
    "Gaji": 3000,
    "LegalName": "Belzar",
    "TanggalLahir": "2024-01-18",
    "TempatLahir": "Left Right Good Night"
}`

func TestCreateSuccess(t *testing.T) {
	body := []byte(inputCreate)
	req := httptest.NewRequest(http.MethodPost, "/api/customer", bytes.NewBuffer(body))
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	customerUsecase.Mock.On("Create", mock.Anything, mock.Anything).Return(&domain.Customer{}, nil)

	customerHandler.Create(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestUpdateSuccess(t *testing.T) {
	body := []byte(inputUpdate)
	req := httptest.NewRequest(http.MethodPut, "/api/customer/:id", bytes.NewBuffer(body))
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	customerUsecase.Mock.On("Update", mock.Anything, mock.Anything).Return(&domain.Customer{}, nil)

	customerHandler.Update(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestFetchByIDSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/customer/:id", nil)
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	customerUsecase.Mock.On("FetchByID", mock.Anything, mock.Anything).Return(domain.Customer{}, nil)

	customerHandler.FetchByID(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestFetchAllSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/customer", nil)
	responseRecorder := httptest.NewRecorder()
	responseRecorder.WriteHeader(http.StatusOK)

	context, _ := gin.CreateTestContext(responseRecorder)
	context.Request = req

	customerUsecase.Mock.On("FetchAll", mock.Anything).Return([]domain.Customer{}, nil)

	customerHandler.FetchAll(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}
