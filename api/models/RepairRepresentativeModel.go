package models

import "time"

type RepairRepresentative struct {
	ID             int64      `json:"id"`
	UserID         *int64     `json:"user_id"`
	Name           string     `json:"name"`
	PhoneNumber    string     `json:"phone_number"`
	Email          string     `json:"email"`
	Active         bool       `json:"active"`
	Certification  string     `json:"certification"`
	Specialization string     `json:"specialization"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
