package usecase

import (
	"context"
	"kredit-plus/domain"
)

type transactionUsecase struct {
	transactionRepo domain.TransactionRepository
}

func NewTransactionUsecase(transactionRepo domain.TransactionRepository) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
	}
}

func (usecase *transactionUsecase) Create(context context.Context, reqBody *domain.TransactionCreateRequest) (*domain.Transaction, error) {
	// Build Transaction struct.
	transactionInput := domain.Transaction{
		CustomerID:     reqBody.CustomerID,
		NomorKontrak:   reqBody.NomorKontrak,
		NomorOnTheRoad: reqBody.NomorOnTheRoad,
		AdminFee:       reqBody.AdminFee,
		JumlahCicilan:  reqBody.JumlahCicilan,
		JumlahBunga:    reqBody.JumlahBunga,
	}

	// Delegate the actual creation to the repository
	return usecase.transactionRepo.Create(context, &transactionInput)
}
