package usecase

import (
	"context"
	"kredit-plus/domain"
)

type customerTenorUsecase struct {
	customerTenorRepo domain.CustomerTenorRepository
	customerRepo      domain.CustomerRepository
}

func NewCustomerTenorUsecase(customerTenorRepo domain.CustomerTenorRepository, customerRepo domain.CustomerRepository) domain.CustomerTenorUsecase {
	return &customerTenorUsecase{
		customerTenorRepo: customerTenorRepo,
		customerRepo:      customerRepo,
	}
}

func (usecase *customerTenorUsecase) Create(context context.Context, reqBody *domain.CustomerTenorCreateRequest) (*domain.CustomerTenor, error) {
	// Check if customer exists by ID.
	customer, err := usecase.customerRepo.FetchByID(context, reqBody.CustomerID)
	if err != nil {
		return nil, err
	}

	// Build Customer Tenor struct.
	customerTenorInput := domain.CustomerTenor{
		CustomerID: customer.ID,
		Bulan:      reqBody.Bulan,
		Limit:      reqBody.Limit,
	}

	// Delegate the actual creation to the repository
	return usecase.customerTenorRepo.Create(context, &customerTenorInput)
}
