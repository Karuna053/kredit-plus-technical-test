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

func (usecase *customerTenorUsecase) Update(context context.Context, reqBody *domain.CustomerTenorUpdateRequest, customerTenorID uint) (*domain.CustomerTenor, error) {
	// Retrieve CustomerTenor.
	customerTenorInput, err := usecase.customerTenorRepo.FetchByID(context, customerTenorID)
	if err != nil {
		return nil, err
	}

	// Update customerTenor struct.
	customerTenorInput.Bulan = reqBody.Bulan
	customerTenorInput.Limit = reqBody.Limit

	// Delegate the actual update process to the repository
	return usecase.customerTenorRepo.Update(context, &customerTenorInput)
}

func (usecase *customerTenorUsecase) FetchByID(context context.Context, customerTenorID uint) (domain.CustomerTenor, error) {
	return usecase.customerTenorRepo.FetchByID(context, customerTenorID)
}

func (usecase *customerTenorUsecase) FetchAll(context context.Context) ([]domain.CustomerTenor, error) {
	return usecase.customerTenorRepo.FetchAll(context)
}

func (usecase *customerTenorUsecase) Delete(context context.Context, customerTenorID uint) error {
	// Retrieve CustomerTenor.
	customerTenorInput, err := usecase.customerTenorRepo.FetchByID(context, customerTenorID)
	if err != nil {
		return err
	}

	return usecase.customerTenorRepo.Delete(context, &customerTenorInput)
}
