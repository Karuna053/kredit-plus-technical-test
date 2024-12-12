package usecase

import (
	"context"
	"kredit-plus/domain"
	"kredit-plus/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var customerTenorRepository = &mocks.CustomerTenorRepository{Mock: mock.Mock{}}
var customerRepository = &mocks.CustomerRepository{Mock: mock.Mock{}}
var testCustomerTenorUsecase = customerTenorUsecase{customerTenorRepository, customerRepository}

func TestCreateSuccess(t *testing.T) {
	customerRepository.Mock.On("FetchByID", context.TODO(), uint(1)).Return(domain.Customer{}, nil)

	customerTenorRepository.Mock.On("Create", context.TODO(), mock.Anything).Return(&domain.CustomerTenor{
		ID:         1,
		CustomerID: 1,
	}, nil)

	customerTenor, err := testCustomerTenorUsecase.Create(context.TODO(), &domain.CustomerTenorCreateRequest{
		CustomerID: 1,
	})

	assert.Nil(t, err)
	assert.NotNil(t, customerTenor)
}

func TestUpdateSuccess(t *testing.T) {
	customerTenorRepository.Mock.On("FetchByID", context.TODO(), uint(1)).Return(domain.CustomerTenor{
		ID:         1,
		CustomerID: 1,
	}, nil)

	customerTenorRepository.Mock.On("Update", context.TODO(), mock.Anything).Return(&domain.CustomerTenor{}, nil)

	customerTenor, err := testCustomerTenorUsecase.Update(context.TODO(), &domain.CustomerTenorUpdateRequest{}, uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, customerTenor)
}

func TestFetchByIDSuccess(t *testing.T) {
	customerTenorRepository.Mock.On("FetchByID", context.TODO(), mock.Anything).Return(domain.CustomerTenor{
		ID:         1,
		CustomerID: 1,
	}, nil)

	customerTenor, err := testCustomerTenorUsecase.FetchByID(context.TODO(), uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, customerTenor)
}

func TestFetchAllSuccess(t *testing.T) {
	customerTenorRepository.Mock.On("FetchAll", context.TODO()).Return([]domain.CustomerTenor{}, nil)

	customerTenors, err := testCustomerTenorUsecase.FetchAll(context.TODO())

	assert.Nil(t, err)
	assert.NotNil(t, customerTenors)
}

func TestDeleteSuccess(t *testing.T) {
	customerTenorRepository.Mock.On("FetchByID", context.TODO(), mock.Anything).Return(domain.CustomerTenor{
		ID:         1,
		CustomerID: 1,
	}, nil)

	customerTenorRepository.Mock.On("Delete", context.TODO(), mock.Anything).Return(nil)

	err := testCustomerTenorUsecase.Delete(context.TODO(), uint(1))

	assert.Nil(t, err)
}
