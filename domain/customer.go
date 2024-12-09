package domain

import "time"

type Customer struct {
	ID         uint      `gorm:"primaryKey"`
	NIK        string    `json:"NIK" gorm:"uniqueIndex"`
	FullName   string    `json:"FullName"`
	LegalName  string    `json:"Legalname"`
	Birthplace string    `json:"Birthplace"`
	Birthday   time.Time `json:"Birthday"`
	Salary     uint      `json:"Salary"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
