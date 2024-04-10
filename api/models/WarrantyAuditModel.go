package models

import (
	"time"
)

type WarrantyAudit struct {
	ID            int64     `json:"id"`
	WarrantyID    int64     `json:"warranty_id"`
	ModifiedField string    `json:"modified_field"`
	OldValue      string    `json:"old_value"`
	NewValue      string    `json:"new_value"`
	ModifiedBy    int64     `json:"modified_by"`
	ModifiedAt    time.Time `json:"modified_at"`
}
