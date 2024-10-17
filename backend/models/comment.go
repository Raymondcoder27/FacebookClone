package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Comment string `gorm:"comment"`
}
