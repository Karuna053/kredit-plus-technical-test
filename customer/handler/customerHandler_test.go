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
