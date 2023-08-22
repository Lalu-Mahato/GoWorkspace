package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"ID" gorm:"primary_key"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
