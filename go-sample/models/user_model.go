package models

import (
	"time"
)

type User struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(username, email, password string) *User {
	currentTime := time.Now()
	return &User{
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}
