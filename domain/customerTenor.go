package domain

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type CustomerTenor struct {
	ID         uint     `gorm:"primaryKey"`
	CustomerID uint     `json:"CustomerID"`
	Bulan      int      `json:"Bulan"`
	Limit      float32  `json:"Limit"`
	Customer   Customer `gorm:"foreignKey:customer_id" json:"-"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CustomerTenorCreateRules struct {
	Bulan int     `validate:"required,min:0,max:12"`
	Limit float32 `validate:"required,min:0,max:1000000000"`
}

type CustomerTenorUpdateRules struct {
	ID    uint    `validate:"required"`
	Bulan int     `validate:"required,min:0,max:12"`
	Limit float32 `validate:"required,min:0,max:1000000000"`
}

type CustomerTenorDeleteRules struct {
	ID uint `validate:"required"`
}
