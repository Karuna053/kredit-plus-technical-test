package repository

import (
	"context"
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
