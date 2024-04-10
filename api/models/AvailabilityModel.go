package models

import "time"

type Availability struct {
	ID                     int64      `json:"id"`
	RepairRepresentativeID int64      `json:"repair_representative_id"`
	AvailabilityType       string     `json:"availability_type"`
	AvailabilityDetails    string     `json:"availability_details"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              *time.Time `json:"updated_at"`
}
