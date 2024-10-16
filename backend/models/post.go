package models

import "gorm.io/gorm"

// Post model
type Post struct {
	gorm.Model
	Content  string    `json:"content"`                           // Replace 'Email' with 'Content' to hold post content
	UserID   uint      `json:"user_id"`                           // Foreign key for User
	User     User      `gorm:"foreignKey:UserID" json:"user"`     // Relationship to User
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments"` // Relationship to Comments
	Image    string    `json:"image"`                             // Add 'Image' field to hold image URL
	Text     string    `json:"text"`                              // Add 'Text' field to hold post text
}
