package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string `json:"password"`
}
