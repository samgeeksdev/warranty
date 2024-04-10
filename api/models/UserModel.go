package models

import (
	"time"
)

type User struct {
	ID           int64      `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password,omitempty"` // Omit password on marshaling
	FirstName    string     `json:"firstName,omitempty"`
	LastName     string     `json:"lastName,omitempty"`
	PhoneNumber  string     `json:"phoneNumber,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
	Active       bool       `json:"active"`
}

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
