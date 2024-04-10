package models

import (
	"time"
)

type Customer struct {
	ID          int64     `json:"id"`
	UserID      *int64    `json:"user_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
}
