package controllers

import (
	"example.com/facebookclone/initializers"
	"example.com/facebookclone/models"
	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	//querry posts and preload related comments and user
	if err := initializers.DB.Preload("Comments").Preload("User").Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"message": "Failed to fetch posts"})
		return
	}
}
