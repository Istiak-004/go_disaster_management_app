package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	ItemName     string  `json:"item_name"`
	Quantity     int     `json:"quantity"`
	PricePerUnit float64 `json:"price_per_unit"`
	PurchasedBy  uint    `json:"purchased_by"`
}
