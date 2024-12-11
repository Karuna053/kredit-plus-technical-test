package domain

import (
	"context"
	"time"

	_ "github.com/go-playground/validator/v10"
)

type CustomerTenor struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint `gorm:"uniqueIndex:idx_customer_bulan"`
	Bulan      int  `gorm:"uniqueIndex:idx_customer_bulan"`
	Limit      float32
	Customer   Customer `gorm:"foreignKey:customer_id" json:"-"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CustomerTenorUsecase interface {
	Create(ctx context.Context, reqBody *CustomerTenorCreateRequest) (*CustomerTenor, error)
	// Update(ctx context.Context, reqBody *CustomerTenorCreateRequest, customerTenorID uint) (*CustomerTenor, error)
}

type CustomerTenorRepository interface {
	Create(ctx context.Context, customerTenorInput *CustomerTenor) (*CustomerTenor, error)
}

// Validation rules for Customer Tenor Create
type CustomerTenorCreateRequest struct {
	CustomerID uint    `json:"CustomerID" validate:"required,min=0,max=4000000000"`
	Bulan      int     `json:"Bulan" validate:"required,min=0,max=12"`
	Limit      float32 `json:"Limit" validate:"required,min=0,max=1000000000"`
}

// Validation rules for Customer Tenor Update
type CustomerTenorUpdateRequest struct {
	CustomerID uint    `json:"CustomerID" validate:"required,min=0,max=4000000000"`
	Bulan      int     `json:"Bulan" validate:"required,min=0,max=12"`
	Limit      float32 `json:"Limit" validate:"required,min=0,max=1000000000"`
}
