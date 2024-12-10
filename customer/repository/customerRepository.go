package repository

import (
	"context"
	"kredit-plus/domain"

	"gorm.io/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{db}
}

func (repo *customerRepository) Create(ctx context.Context, customerInput *domain.Customer) (*domain.Customer, error) {
	err := repo.DB.Create(&customerInput).Error
	if err != nil {
		return nil, err // Return nil for the customer and the error
	}

	return customerInput, nil // Return the inserted customer and nil error
}
