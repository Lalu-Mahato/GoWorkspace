package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `json:"id" gorm:"primary_key"`
	Name      string         `json:"name" gorm:"size:255" validate:"required"`
	Email     string         `json:"email" gorm:"size:255" validate:"required,email"`
	Password  string         `json:"password" gorm:"size:255" validate:"required,min=6"`
	Age       int            `json:"age" validate:"gte=18"`
	Status    bool           `json:"status" gorm:"default:true"`
	CreatedAt time.Time      `json:"deleteddAt" gorm:"index"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uint64(time.Now().UnixNano())
	return
}
