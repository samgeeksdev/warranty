package models

import (
	"time"
)

type Product struct {
	ID             int64      `json:"id"`
	Name           string     `json:"name"`
	Slug           string     `json:"slug"`
	Description    string     `json:"description"`
	CategoryID     int        `json:"category_id"`
	ManufacturerID *int       `json:"manufacturer_id"`
	Model          string     `json:"model"`
	SKU            string     `json:"sku"`
	Price          float64    `json:"price"`
	Stock          int        `json:"stock"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	WarrantyMonths *int       `json:"warranty_months"`
	Retired        bool       `json:"retired"`
}
