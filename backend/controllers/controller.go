package controllers

import (
	"example.com/facebookclone/models"
	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	//querry posts and preload related comments and user
	if err := 
}
