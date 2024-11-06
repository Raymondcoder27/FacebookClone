package controllers

import (
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
			Image:     string(post.ImageData),
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

	currentTime := time.Now()

	// c.JSON(200, gin.H{"posts": postResponses})
	c.IndentedJSON(http.StatusOK, gin.H{"code": 200, "data": postResponses, "timestamp": currentTime})
}

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

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Fetch the user from your database (e.g., based on session or JWT token)
	var user models.User
	if err := initializers.DB.Where("id = ?", userID).First(&user).Error; err != nil { // Replace `userID` appropriately
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	// Update the image using the service
	updatedUser, err := services.UpdateUserImage(&user, fileHeader, width, height, left, top)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save the updated user model to the database
	if err := initializers.DB.Save(updatedUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user image"})
		return
	}

	// Return the updated user with the new image
	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}

// GetPosts returns a list of all posts with user and comments
// func GetPosts(c *gin.Context) {
// 	var posts []models.Post

// 	// Fetch posts with associated User and Comments, order by created_at descending
// 	err := initializers.DB.Preload("User").Preload("Comments.User").Order("created_at desc").Find(&posts).Error
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve posts"})
// 		return
// 	}

// 	// Custom response formatting if needed
// 	c.JSON(http.StatusOK, gin.H{"posts": posts})
// }

func GetPosts(c *gin.Context) {
	var posts []models.Post

	// Fetch posts with associated User and Comments, order by created_at descending
	err := initializers.DB.Preload("User").Preload("Comments.User").Order("created_at desc").Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve posts"})
		return
	}

	// Transform posts with base64-encoded images for response
	var postResponses []map[string]interface{}
	for _, post := range posts {
		// Encode the image data to base64
		imageBase64 := base64.StdEncoding.EncodeToString(post.ImageData)

		// Create a post response map
		postResponse := map[string]interface{}{
			"id":         post.ID,
			"text":       post.Text,
			"image":      imageBase64,
			"created_at": post.CreatedAt.Format("2006-01-02 15:04:05"),
			"user": map[string]interface{}{
				"id":    post.User.ID,
				"name":  post.User.Name,
				"email": post.User.Email,
			},
			"comments": post.Comments, // Adapt if you want custom formatting for comments
		}

		postResponses = append(postResponses, postResponse)
	}

	// Return the custom response format
	c.JSON(http.StatusOK, gin.H{"posts": postResponses})
}

// // CreatePost handles creating a new post
// func CreatePost(c *gin.Context) {
// 	// Validate the text field
// 	text := c.PostForm("text")
// 	image := c.PostForm("image")
// 	if text == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Text is required"})
// 		return
// 	}

// 	if image == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
// 		return
// 	}

// 	// Initialize a new post
// 	var post models.Post

// 	// Handle image upload if provided
// 	fileHeader, err := c.FormFile("image")
// 	if err == nil && fileHeader != nil {
// 		width, _ := strconv.Atoi(c.PostForm("width"))
// 		height, _ := strconv.Atoi(c.PostForm("height"))
// 		left, _ := strconv.Atoi(c.PostForm("left"))
// 		top, _ := strconv.Atoi(c.PostForm("top"))

// 		updatedPost, err := services.UpdateImage(&post, fileHeader, width, height, left, top)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Image upload failed"})
// 			return
// 		}
// 		post = *updatedPost
// 	}

// 	// Assume the user is authenticated, and their ID is available in the context
// 	userID := c.MustGet("userID").(uint) // This depends on how you've set up authentication

// 	post.UserID = userID
// 	post.Text = text
// 	post.Image = image

// 	// Save the post to the database
// 	if err := initializers.DB.Create(&post).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save post"})
// 		return
// 	}

// 	// Return success response
// 	c.JSON(http.StatusOK, gin.H{"post": post})
// }

func CreatePost(c *gin.Context) {
	// Get the text field
	text := c.PostForm("text")
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text is required"})
		return
	}

	// fileHeader, err := c.FormFile("image")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
	// 	return
	// }

	// Handle image upload if provided
	// var imagePath string
	// fileHeader, err := c.FormFile("image")
	// if err == nil {
	// 	imagePath = "./public/" // Define the path where you want to save it
	// 	if err := c.SaveUploadedFile(fileHeader, imagePath); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Image upload failed"})
	// 		return
	// 	}
	// }
	// Open the image file

	// file, err := fileHeader.Open()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open image file"})
	// 	return
	// }
	// defer file.Close()

	// Read the image file data into a byte slice
	// imageData := make([]byte, fileHeader.Size)
	// _, err = file.Read(imageData)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read image data"})
	// 	return
	// }

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	// Populate post fields
	var post models.Post
	post.UserID = userID.(uint) // Adjust for your authentication setup
	post.Text = text
	// post.ImageData = imageData
	// post.Image = imagePath // Store the file path

	// Save the post to the database
	if err := initializers.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// DeletePost deletes a post by its ID
func DeletePost(c *gin.Context) {
	// Get the post ID from the URL parameter
	postID := c.Param("id")
	var post models.Post

	// Find the post
	if err := initializers.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if the post has an associated image and delete it
	if string(post.ImageData) != "" {
		imagePath := filepath.Join("public", filepath.Base(string(post.ImageData)))
		if _, err := os.Stat(imagePath); err == nil {
			if err := os.Remove(imagePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete image"})
				return
			}
		}
	}

	// Delete the post
	if err := initializers.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete post"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func CommentStore(c *gin.Context) {
	// Get request data
	var requestBody struct {
		Text   string `json:"text" binding:"required"`
		PostID uint   `json:"post_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from context (set by authentication middleware)
	user, _ := c.Get("user")

	// Create the comment
	comment := models.Comment{
		Text:   requestBody.Text,
		UserID: user.(models.User).ID,
		PostID: requestBody.PostID,
	}

	// Save to the database
	if err := initializers.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully", "comment": comment})
}

// Destroy a comment by ID
func CommentDestroy(c *gin.Context) {
	// Get comment ID from the URL
	id := c.Param("id")

	// Find the comment
	var comment models.Comment
	if err := initializers.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Delete the comment
	if err := initializers.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// Show posts of the logged-in user
func UserIndex(c *gin.Context) {
	// Get user from context (set by authentication middleware)
	user, _ := c.Get("user")

	// Find posts of the user
	var posts []models.Post
	if err := initializers.DB.Where("user_id = ?", user.(models.User).ID).Order("created_at desc").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// Show a specific user by ID along with their posts
func UserShow(c *gin.Context) {
	// Get user ID from the URL
	id := c.Param("id")

	// Find user by ID
	var user models.User
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Find posts of the user
	var posts []models.Post
	if err := initializers.DB.Where("user_id = ?", id).Order("created_at desc").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"posts": posts,
	})
}

// Update user profile image
// func UserUpdateImage(c *gin.Context) {
// 	// Get user from context (set by authentication middleware)
// 	user, _ := c.Get("user")

// 	// Validate image file
// 	file, err := c.FormFile("image")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
// 		return
// 	}

// 	// Use the ImageService to update the image
// 	if err := services.UpdateImage(user.(models.User), file, c); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update image"})
// 		return
// 	}

// 	// Save updated user
// 	if err := initializers.DB.Save(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user image"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Profile image updated successfully", "user": user})
// }

// get user
func GetLoggedInUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
