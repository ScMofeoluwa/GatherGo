package entity

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Verified     bool      `json:"verified"`
	RegisteredAt time.Time `json:"registeredAt"`
}

type CreateUser struct {
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}
