package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint64         `json:"id" gorm:"primary_key"`
	Name      string         `json:"name" gorm:"size:255"`
	Content   string         `json:"content" gorm:"type:text"`
	CreatedAt time.Time      `json:"deleteddAt" gorm:"index"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
