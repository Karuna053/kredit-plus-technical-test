package domain

import (
	"context"
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID             uint `gorm:"primaryKey"`
	CustomerID     uint
	NomorKontrak   string
	NomorOnTheRoad string
	AdminFee       int
	JumlahCicilan  float32
	JumlahBunga    float32
	Customer       Customer `gorm:"foreignKey:customer_id" json:"-"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TransactionUsecase interface {
	Create(ctx context.Context, reqBody *TransactionCreateRequest) (*Transaction, error)
	// Update(ctx context.Context, reqBody *TransactionUpdateRequest, transactionID uint) (*Transaction, error)
	// FetchByID(ctx context.Context, transactionID uint) (Customer, error)
	// FetchAll(ctx context.Context) ([]Transaction, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, transactionInput *Transaction) (*Transaction, error)
	// Update(ctx context.Context, transactionInput *Transaction) (*Transaction, error)
	// FetchByID(ctx context.Context, transactionID uint) (Transaction, error)
	// FetchAll(ctx context.Context) ([]Customer, error)
}

// Validation rules for Transaction Create
type TransactionCreateRequest struct {
	CustomerID     uint    `json:"CustomerID" validate:"required,min=0,max=4000000000"`
	NomorKontrak   string  `json:"NomorKontrak" validate:"required,max=255"`
	NomorOnTheRoad string  `json:"NomorOnTheRoad" validate:"required,max=255"`
	AdminFee       int     `json:"AdminFee" validate:"required,min=0,max=1000000000"`
	JumlahCicilan  float32 `json:"JumlahCicilan" validate:"required,min=0,max=1000000000"`
	JumlahBunga    float32 `json:"JumlahBunga" validate:"required,min=0,max=1000000000"`
}

// Validation rules for Transaction Update
type TransactionUpdateRequest struct {
	CustomerID     uint    `json:"CustomerID" validate:"required,min=0,max=4000000000"`
	NomorKontrak   string  `json:"NomorKontrak" validate:"required,max=255"`
	NomorOnTheRoad string  `json:"NomorOnTheRoad" validate:"required,max=255"`
	AdminFee       int     `json:"AdminFee" validate:"required,min=0,max=1000000000"`
	JumlahCicilan  float32 `json:"JumlahCicilan" validate:"required,min=0,max=1000000000"`
	JumlahBunga    float32 `json:"JumlahBunga" validate:"required,min=0,max=1000000000"`
}
