package usecase

import (
	"context"
	"kredit-plus/domain"
	"kredit-plus/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transactionRepository = &mocks.TransactionRepository{Mock: mock.Mock{}}
var testTransactionUsecase = transactionUsecase{transactionRepository}

func TestCreateSuccess(t *testing.T) {
	transactionRepository.Mock.On("Create", context.TODO(), mock.Anything).Return(&domain.Transaction{}, nil)

	transaction, err := testTransactionUsecase.Create(context.TODO(), &domain.TransactionCreateRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}

func TestUpdateSuccess(t *testing.T) {
	transactionRepository.Mock.On("FetchByID", context.TODO(), uint(1)).Return(domain.Transaction{}, nil)
	transactionRepository.Mock.On("Update", context.TODO(), mock.Anything).Return(&domain.Transaction{}, nil)

	transaction, err := testTransactionUsecase.Update(context.TODO(), &domain.TransactionUpdateRequest{}, uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}

func TestFetchByIDSuccess(t *testing.T) {
	transactionRepository.Mock.On("FetchByID", context.TODO(), mock.Anything).Return(domain.Transaction{}, nil)

	transaction, err := testTransactionUsecase.FetchByID(context.TODO(), uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}

func TestFetchAllSuccess(t *testing.T) {
	transactionRepository.Mock.On("FetchAll", context.TODO()).Return([]domain.Transaction{}, nil)

	transactions, err := testTransactionUsecase.FetchAll(context.TODO())

	assert.Nil(t, err)
	assert.NotNil(t, transactions)
}
