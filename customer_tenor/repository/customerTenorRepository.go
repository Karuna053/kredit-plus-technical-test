package repository

import (
	"context"
	"kredit-plus/domain"

	"gorm.io/gorm"
)

type customerTenorRepository struct {
	DB *gorm.DB
}

func NewCustomerTenorRepository(db *gorm.DB) domain.CustomerTenorRepository {
	return &customerTenorRepository{db}
}

func (repo *customerTenorRepository) Create(ctx context.Context, customerTenorInput *domain.CustomerTenor) (*domain.CustomerTenor, error) {
	err := repo.DB.Create(&customerTenorInput).Error
	if err != nil {
		return nil, err
	}

	return customerTenorInput, nil // Return the inserted customer and nil error
}
