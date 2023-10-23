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
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}
