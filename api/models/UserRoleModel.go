package models

import "time"

type UserRole struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}
