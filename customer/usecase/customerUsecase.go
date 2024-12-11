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
	parsedBirthdate, err := time.Parse("2006-01-02", reqBody.TanggalLahir)
	if err != nil {
		return nil, err // Return nil for the customer and the error
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
