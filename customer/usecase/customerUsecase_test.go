package usecase

import (
	"context"
	"kredit-plus/domain"
	"kredit-plus/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var customerRepository = &mocks.CustomerRepository{Mock: mock.Mock{}}
var customerUseCase = customerUsecase{customerRepository}

func TestCreateSuccess(t *testing.T) {
	customerRepository.Mock.On("Create", context.TODO(), mock.Anything).Return(&domain.Customer{}, nil)

	customer, err := customerUseCase.Create(context.TODO(), &domain.CustomerCreateRequest{
		NIK:          "123456789",
		FullName:     "Eilang Whites",
		LegalName:    "Testing 123",
		TempatLahir:  "Jakarta",
		TanggalLahir: "2013-08-07",
		Gaji:         200000,
		FotoKTP:      "http://google.com",
		FotoSelfie:   "http://googlefu.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestUpdateSuccess(t *testing.T) {
	customerRepository.Mock.On("FetchByID", context.TODO(), uint(1)).Return(domain.Customer{
		NIK:          "00000000",
		FullName:     "Truman",
		LegalName:    "Helvetrr",
		TempatLahir:  "Bangkula",
		TanggalLahir: time.Now(),
		Gaji:         200000,
		FotoKTP:      "http://google.com",
		FotoSelfie:   "http://googlefu.com",
	}, nil)

	customerRepository.Mock.On("Update", context.TODO(), mock.Anything).Return(&domain.Customer{}, nil)

	customer, err := customerUseCase.Update(context.TODO(), &domain.CustomerUpdateRequest{
		NIK:          "00000000",
		FullName:     "Eilang Whites",
		LegalName:    "Testing 123",
		TempatLahir:  "Jakarta",
		TanggalLahir: "2013-08-07",
		Gaji:         200000,
		FotoKTP:      "http://google.com",
		FotoSelfie:   "http://googlefu.com",
	}, uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestFetchByIDSuccess(t *testing.T) {
	customerRepository.Mock.On("FetchByID", context.TODO(), uint(1)).Return(domain.Customer{
		NIK:          "00000000",
		FullName:     "Truman",
		LegalName:    "Helvetrr",
		TempatLahir:  "Bangkula",
		TanggalLahir: time.Now(),
		Gaji:         200000,
		FotoKTP:      "http://google.com",
		FotoSelfie:   "http://googlefu.com",
	}, nil)

	customer, err := customerUseCase.FetchByID(context.TODO(), uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestFetchAllSuccess(t *testing.T) {
	customerRepository.Mock.On("FetchAll", context.TODO()).Return([]domain.Customer{}, nil)

	customers, err := customerUseCase.FetchAll(context.TODO())

	assert.Nil(t, err)
	assert.NotNil(t, customers)
}
