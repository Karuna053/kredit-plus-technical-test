package usecase

import (
	"context"
	"kredit-plus/domain"
	"time"
)

type customerUsecase struct {
	customerRepo domain.CustomerRepository
}

func NewCustomerUsecase(customerRepo domain.CustomerRepository) domain.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (usecase *customerUsecase) Create(context context.Context, reqBody *domain.CustomerCreateRequest) (*domain.Customer, error) {
	// Validate birthdate.
	parsedBirthdate, err := time.Parse("2006-01-02", reqBody.TanggalLahir)
	if err != nil {
		return nil, err
	}

	// Build Customer struct.
	customerInput := domain.Customer{
		NIK:          reqBody.NIK,
		FullName:     reqBody.FullName,
		LegalName:    reqBody.LegalName,
		TempatLahir:  reqBody.TempatLahir,
		TanggalLahir: parsedBirthdate,
		Gaji:         reqBody.Gaji,
		FotoKTP:      reqBody.FotoKTP,
		FotoSelfie:   reqBody.FotoSelfie,
	}

	// Delegate the actual creation to the repository
	return usecase.customerRepo.Create(context, &customerInput)
}

func (usecase *customerUsecase) Update(context context.Context, reqBody *domain.CustomerUpdateRequest, customerID uint) (*domain.Customer, error) {
	// Validate birthdate.
	parsedBirthdate, err := time.Parse("2006-01-02", reqBody.TanggalLahir)
	if err != nil {
		return nil, err
	}

	// Retrieve Customer.
	customerInput, err := usecase.customerRepo.FetchByID(context, customerID)
	if err != nil {
		return nil, err
	}

	// Update customer struct.
	customerInput.NIK = reqBody.NIK
	customerInput.FullName = reqBody.FullName
	customerInput.LegalName = reqBody.LegalName
	customerInput.TempatLahir = reqBody.TempatLahir
	customerInput.TanggalLahir = parsedBirthdate
	customerInput.Gaji = reqBody.Gaji
	customerInput.FotoKTP = reqBody.FotoKTP
	customerInput.FotoSelfie = reqBody.FotoSelfie

	// Delegate the actual update process to the repository
	return usecase.customerRepo.Update(context, &customerInput)
}
