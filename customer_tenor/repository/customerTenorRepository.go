package repository

import (
	"context"
	"errors"
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

func (repo *customerTenorRepository) Update(ctx context.Context, customerTenorInput *domain.CustomerTenor) (*domain.CustomerTenor, error) {
	err := repo.DB.Save(&customerTenorInput).Error
	if err != nil {
		return nil, err
	}

	return customerTenorInput, nil // Return the updated customerTenor and nil error
}

func (repo *customerTenorRepository) FetchByID(ctx context.Context, customerTenorID uint) (domain.CustomerTenor, error) {
	var customerTenor domain.CustomerTenor

	err := repo.DB.First(&customerTenor, "id = ?", customerTenorID).Error
	if err != nil {
		return domain.CustomerTenor{}, errors.New("customerTenor ID does not exist in database")
	}

	return customerTenor, nil
}

func (repo *customerTenorRepository) FetchAll(ctx context.Context) ([]domain.CustomerTenor, error) {
	var customerTenors []domain.CustomerTenor

	err := repo.DB.Find(&customerTenors).Error
	if err != nil {
		return []domain.CustomerTenor{}, err
	}

	return customerTenors, nil
}

func (repo *customerTenorRepository) Delete(context context.Context, customerTenorInput *domain.CustomerTenor) error {
	err := repo.DB.Delete(&customerTenorInput).Error
	if err != nil {
		return err
	}

	return err
}
