package domain

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID             uint     `gorm:"primaryKey"`
	CustomerID     uint     `json:"CustomerID"`
	NomorKontrak   string   `json:"NomorKontrak"`
	NomorOnTheRoad string   `json:"NomorOnTheRoad"`
	AdminFee       uint     `json:"AdminFee"`
	JumlahCicilan  float32  `json:"JumlahCicilan"`
	JumlahBunga    float32  `json:"JumlahBunga"`
	Customer       Customer `gorm:"foreignKey:customer_id" json:"-"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
