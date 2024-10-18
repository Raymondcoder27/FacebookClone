package controllers

import (
	"net/http"
	"strconv"

	"example.com/facebookclone/contracts"
	"example.com/facebookclone/initializers"
	"example.com/facebookclone/models"
	"example.com/facebookclone/services"
	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	//querry posts and preload related comments and user
	if err := initializers.DB.Preload("Comments").Preload("User").Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"message": "Failed to fetch posts"})
		return
	}

	//map posts to the post response format
	var postResponses []contracts.PostResponse

	for _, post := range posts {
		postResponse := contracts.PostResponse{
			ID:        post.ID,
			Text:      post.Text,
			Image:     post.Image,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			User: contracts.UserResponse{
				ID:    post.User.ID,
				Name:  post.User.Name,
				Email: post.User.Email,
				Image: "",
			},
		}

		//map comments to the comment response format
		var commentsResponse []contracts.CommentResponse

		for _, comment := range post.Comments {
			commentResponse := contracts.CommentResponse{
				ID:   comment.ID,
				Text: comment.Text,
				User: contracts.UserResponse{
					ID:    comment.User.ID,
					Name:  comment.User.Name,
					Email: comment.User.Email,
					Image: "",
				},
			}
			commentsResponse = append(commentsResponse, commentResponse)
		}
		postResponses = append(postResponses, postResponse)
	}

	c.JSON(200, gin.H{"posts": postResponses})
}

// UpdateUserImage updates the user's profile picture
func UpdateUserImage(c *gin.Context) {
	// Get the uploaded file from the request
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload"})
		return
	}

	// Get cropping dimensions from the form
	width, _ := strconv.Atoi(c.PostForm("width"))
	height, _ := strconv.Atoi(c.PostForm("height"))
	left, _ := strconv.Atoi(c.PostForm("left"))
	top, _ := strconv.Atoi(c.PostForm("top"))

	// Fetch the user from your database (e.g., based on session or JWT token)
	var user models.User // In a real case, you would get the user from DB
	// Simulating fetching user. Replace with DB logic:
	user.ID = 1
	user.Image = "/images/old_image.png"

	// Update the image using the service
	updatedUser, err := services.UpdateUserImage(&user, fileHeader, width, height, left, top)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated user with the new image
	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}

// GetPosts returns a list of all posts with user and comments
func GetPosts(c *gin.Context) {
	var posts []models.Post

	// Fetch posts with associated User and Comments, order by created_at descending
	err := initializers.DB.Preload("User").Preload("Comments.User").Order("created_at desc").Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve posts"})
		return
	}

	// Custom response formatting if needed
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// CreatePost handles creating a new post
func CreatePost(c *gin.Context) {
	// Validate the text field
	text := c.PostForm("text")
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text is required"})
		return
	}

	// Initialize a new post
	var post models.Post

	// Handle image upload if provided
	fileHeader, err := c.FormFile("image")
	if err == nil && fileHeader != nil {
		width, _ := strconv.Atoi(c.PostForm("width"))
		height, _ := strconv.Atoi(c.PostForm("height"))
		left, _ := strconv.Atoi(c.PostForm("left"))
		top, _ := strconv.Atoi(c.PostForm("top"))

		updatedPost, err := services.UpdateImage(&post, fileHeader, width, height, left, top)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Image upload failed"})
			return
		}
		post = *updatedPost
	}

	// Assume the user is authenticated, and their ID is available in the context
	userID := c.MustGet("userID").(uint) // This depends on how you've set up authentication

	post.UserID = userID
	post.Text = text

	// Save the post to the database
	if err := models.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save post"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"post": post})
}
