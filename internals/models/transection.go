package models

import "gorm.io/gorm"

type Transaction struct {
    gorm.Model
    Amount       float64 `json:"amount"`
    TransactionType string `json:"transaction_type"` 
}