package repository

import (
	"context"
	"errors"
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
		return nil, err
	}

	return customerInput, nil // Return the inserted customer and nil error
}

func (repo *customerRepository) Update(ctx context.Context, customerInput *domain.Customer) (*domain.Customer, error) {
	err := repo.DB.Save(&customerInput).Error
	if err != nil {
		return nil, err
	}

	return customerInput, nil // Return the updated customer and nil error
}

func (repo *customerRepository) FetchByID(ctx context.Context, customerID uint) (domain.Customer, error) {
	var customer domain.Customer

	err := repo.DB.First(&customer, "id = ?", customerID).Error
	if err != nil {
		return domain.Customer{}, errors.New("customer ID does not exist in database")
	}

	return customer, nil
}

func (repo *customerRepository) FetchAll(ctx context.Context) ([]domain.Customer, error) {
	var customers []domain.Customer

	err := repo.DB.Find(&customers).Error
	if err != nil {
		return []domain.Customer{}, err
	}

	return customers, nil
}
