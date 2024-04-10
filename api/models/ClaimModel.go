package models

import "time"

type Claim struct {
	ID          int64     `json:"id"`
	WarrantyID  int64     `json:"warranty_id"`
	Description string    `json:"description"`
	ClaimDate   time.Time `json:"claim_date"`
	StatusID    int       `json:"status_id"`
	Resolution  string    `json:"resolution"`
	ClaimFiles  string    `json:"claim_files_json"`
}
