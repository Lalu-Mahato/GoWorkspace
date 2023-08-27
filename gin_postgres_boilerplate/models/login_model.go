package models

type Login struct {
	Email    string `json:"email" gorm:"size:255" validate:"required,email"`
	Password string `json:"password" gorm:"size:255" validate:"required,min=6"`
}
