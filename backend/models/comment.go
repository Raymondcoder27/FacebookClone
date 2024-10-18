package models

import "gorm.io/gorm"

// Comment model
type Comment struct {
	gorm.Model
	Comment string `gorm:"comment" json:"comment"`        // Add json tag for marshaling
	PostID  uint   `json:"post_id"`                       // Foreign key for Post
	UserID  uint   `json:"user_id"`                       // Foreign key for User
	Post    Post   `gorm:"foreignKey:PostID" json:"post"` // Relationship to Post
	User    User   `gorm:"foreignKey:UserID" json:"user"` // Relationship to User
	Text    string `json:"text"`                          // Add 'Text' field to hold comment text
}
