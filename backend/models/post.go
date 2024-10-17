package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string `json:"password"`
}
