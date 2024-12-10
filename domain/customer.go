package domain

import (
	"context"
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Customer struct {
	ID           uint   `gorm:"primaryKey"`
	NIK          string `gorm:"uniqueIndex"`
	FullName     string
	LegalName    string
	TempatLahir  string
	TanggalLahir time.Time
	Gaji         float32
	FotoKTP      string
	FotoSelfie   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CustomerUsecase interface {
	Create(ctx context.Context, reqBody *CustomerCreateRequest) (*Customer, error)
}

type CustomerRepository interface {
	Create(ctx context.Context, customerInput *Customer) (*Customer, error)
}

// Validation rules for Customer Create
type CustomerCreateRequest struct {
	NIK          string    `json:"NIK" validate:"required,max=255"`
	FullName     string    `json:"FullName" validate:"required,max=255"`
	LegalName    string    `json:"LegalName" validate:"required,max=255"`
	TempatLahir  string    `json:"TempatLahir" validate:"required,max=255"`
	TanggalLahir time.Time `json:"TanggalLahir" validate:"required"`
	Gaji         float32   `json:"Gaji" validate:"required,max=1000000000"`
	FotoKTP      string    `json:"FotoKTP" validate:"required,max=2000"`
	FotoSelfie   string    `json:"FotoSelfie" validate:"required,max=2000"`
}

// Validation rules for Customer Update
type CustomerUpdateRequest struct {
	ID           uint      `json:"ID" validate:"required"`
	NIK          string    `json:"NIK" validate:"required,max=255"`
	FullName     string    `json:"FullName" validate:"required,max=255"`
	LegalName    string    `json:"LegalName" validate:"required,max=255"`
	TempatLahir  string    `json:"TempatLahir" validate:"required,max=255"`
	TanggalLahir time.Time `json:"TanggalLahir" validate:"required"`
	Gaji         float32   `json:"Gaji" validate:"required,min=0,max=1000000000"`
	FotoKTP      string    `json:"FotoKTP" validate:"required,max=2000"`
	FotoSelfie   string    `json:"FotoSelfie" validate:"required,max=2000"`
}
