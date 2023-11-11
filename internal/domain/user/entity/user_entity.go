package entity

import "time"

type User struct {
	RegisteredAt time.Time `json:"registeredAt"`
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Verified     bool      `json:"verified"`
}

type CreateUser struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
