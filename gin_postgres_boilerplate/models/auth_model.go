package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Login struct {
	Email    string `json:"email" gorm:"size:255" validate:"required,email"`
	Password string `json:"password" gorm:"size:255" validate:"required,min=6"`
}

type LoggedInResponse struct {
	ID        uint64 `json:"ID"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    bool   `json:"status" `
	Token     string `json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type JWTClaim struct {
	ID    uint64 `json:"ID"`
	Email string `json:"email"`
	jwt.StandardClaims
}
