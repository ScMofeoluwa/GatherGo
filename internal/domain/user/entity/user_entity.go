package entity

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Verified     bool      `json:"verified"`
	RegisteredAt time.Time `json:"registeredAt"`
}
