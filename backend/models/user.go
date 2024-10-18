package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `gorm:"unique" json:"email"`
	Password string    `json:"password"`
	Posts    []Post    `gorm:"foreignKey:UserID" json:"posts"`    // Relationship to Posts
	Comments []Comment `gorm:"foreignKey:UserID" json:"comments"` // Relationship to Comments
	Image    string    `json:"image"`                             // Add 'Image' field to hold image URL
}
