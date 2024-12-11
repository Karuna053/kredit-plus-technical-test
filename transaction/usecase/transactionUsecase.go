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

func (usecase *transactionUsecase) Update(context context.Context, reqBody *domain.TransactionUpdateRequest, transactionID uint) (*domain.Transaction, error) {
	// Retrieve Transaction.
	transactionInput, err := usecase.transactionRepo.FetchByID(context, transactionID)
	if err != nil {
		return nil, err
	}

	// Update transaction struct.
	transactionInput.CustomerID = reqBody.CustomerID
	transactionInput.NomorKontrak = reqBody.NomorKontrak
	transactionInput.NomorOnTheRoad = reqBody.NomorOnTheRoad
	transactionInput.AdminFee = reqBody.AdminFee
	transactionInput.JumlahCicilan = reqBody.JumlahCicilan
	transactionInput.JumlahBunga = reqBody.JumlahBunga

	// Delegate the actual update process to the repository
	return usecase.transactionRepo.Update(context, &transactionInput)
}

func (usecase *transactionUsecase) FetchByID(context context.Context, transactionID uint) (domain.Transaction, error) {
	return usecase.transactionRepo.FetchByID(context, transactionID)
}

func (usecase *transactionUsecase) FetchAll(context context.Context) ([]domain.Transaction, error) {
	return usecase.transactionRepo.FetchAll(context)
}
