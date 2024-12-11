package repository

import (
	"context"
	"errors"
	"kredit-plus/domain"

	"gorm.io/gorm"
)

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{db}
}

func (repo *transactionRepository) Create(ctx context.Context, transactionInput *domain.Transaction) (*domain.Transaction, error) {
	err := repo.DB.Create(&transactionInput).Error
	if err != nil {
		return nil, err
	}

	return transactionInput, nil // Return the inserted transaction and nil error
}

func (repo *transactionRepository) Update(ctx context.Context, transactionInput *domain.Transaction) (*domain.Transaction, error) {
	err := repo.DB.Save(&transactionInput).Error
	if err != nil {
		return nil, err
	}

	return transactionInput, nil // Return the updated transaction and nil error
}

func (repo *transactionRepository) FetchByID(ctx context.Context, transactionID uint) (domain.Transaction, error) {
	var transaction domain.Transaction

	err := repo.DB.First(&transaction, "id = ?", transactionID).Error
	if err != nil {
		return domain.Transaction{}, errors.New("transaction ID does not exist in database")
	}

	return transaction, nil
}

func (repo *transactionRepository) FetchAll(ctx context.Context) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	err := repo.DB.Find(&transactions).Error
	if err != nil {
		return []domain.Transaction{}, err
	}

	return transactions, nil
}
