package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64         `json:"ID" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:255" validate:"required"`
	Email     string         `json:"email" gorm:"size:255" validate:"required,email"`
	Password  string         `json:"password" gorm:"size:255" validate:"required,min=6"`
	Status    bool           `json:"status" gorm:"default:true"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserWithoutPassword struct {
	ID        uint64 `json:"ID"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    bool   `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uint64(time.Now().UnixNano())
	return
}
