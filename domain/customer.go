package domain

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Customer struct {
	ID           uint      `gorm:"primaryKey"`
	NIK          string    `json:"NIK" gorm:"uniqueIndex"`
	FullName     string    `json:"FullName"`
	LegalName    string    `json:"Legalname"`
	TempatLahir  string    `json:"TempatLahir"`
	TanggalLahir time.Time `json:"TanggalLahir"`
	Gaji         float32   `json:"Gaji"`
	FotoKTP      string    `json:"FotoKTP"`
	FotoSelfie   string    `json:"FotoSelfie"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CustomerCreateRules struct { // Used in create context.
	NIK          string    `validate:"required,max=255"`
	FullName     string    `validate:"required,max=255"`
	LegalName    string    `validate:"required,max=255"`
	TempatLahir  string    `validate:"required,max=255"`
	TanggalLahir time.Time `validate:"required"`
	Gaji         float32   `validate:"required,max=1000000000"`
	FotoKTP      string    `validate:"required,max=2000"`
	FotoSelfie   string    `validate:"required,max=2000"`
}

type CustomerUpdateRules struct { // Used in update context.
	ID           uint      `validate:"required"`
	NIK          string    `validate:"required,max=255"`
	FullName     string    `validate:"required,max=255"`
	LegalName    string    `validate:"required,max=255"`
	TempatLahir  string    `validate:"required,max=255"`
	TanggalLahir time.Time `validate:"required"`
	Gaji         float32   `validate:"required,min=0,max=1000000000"`
	FotoKTP      string    `validate:"required,max=2000"`
	FotoSelfie   string    `validate:"required,max=2000"`
}
