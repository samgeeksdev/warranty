package models

import (
	"time"
)

type Warranty struct {
	ID             int64      `json:"id"`
	ProductID      int64      `json:"product_id"`
	WarrantyTypeID int        `json:"warranty_type_id"`
	DurationMonths int        `json:"duration_months"`
	StartDate      time.Time  `json:"start_date"`
	EndDate        *time.Time `json:"end_date"`
	RegisteredBy   int64      `json:"registered_by"`
	CustomerID     *int64     `json:"customer_id"`
}
