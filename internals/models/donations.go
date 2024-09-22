package models

import (
	"gorm.io/gorm"
)

type Donation struct {
	gorm.Model
	Amount    float64 `json:"amount"`
	Message   string  `json:"message"`
	DonorName string  `json:"donor_name"`
}

type Expense struct {
	gorm.Model
	Amount    float64 `json:"amount"`
	ItemName  string  `json:"item_name"`
	Message   string  `json:"message"`
	ExpenseBy string  `json:"expense_by"`
}
