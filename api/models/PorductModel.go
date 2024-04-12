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

type ProductFilter struct {
	CategoryID *uint   `json:"category_id,omitempty"`
	Name       *string `json:"name,omitempty"` // Add name filter field
	// You can add other filter fields here (e.g., price range, manufacturer ID)
}
