package main

import (
	"example.com/facebookclone/auth"
	"example.com/facebookclone/controllers"
	"example.com/facebookclone/initializers"
	"example.com/facebookclone/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateDB()
}

func main() {
	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	r.Use(cors.New(config))

	r.POST("/register", auth.Signup)
	r.POST("/login", auth.Login)
	r.GET("/validate", middleware.RequireAuth, auth.Validate)

	// Protected routes (middleware required)
	authGroup := r.Group("/", middleware.RequireAuth)
	{
		authGroup.GET("/", controllers.GetAllPosts)           // Home route to get posts
		authGroup.POST("/post", controllers.CreatePost)       // Route to create a post
		authGroup.DELETE("/post/:id", controllers.DeletePost) // Route to delete a post

		authGroup.POST("/comment", controllers.CreateComment)       // Route to create a comment
		authGroup.DELETE("/comment/:id", controllers.DeleteComment) // Route to delete a comment

		authGroup.GET("/user/:id", controllers.GetUser)                   // Route to get user details
		authGroup.POST("/user/update-image", controllers.UpdateUserImage) // Route to update user image

		authGroup.GET("/profile", controllers.EditProfile)      // Route to edit profile
		authGroup.PATCH("/profile", controllers.UpdateProfile)  // Route to update profile
		authGroup.DELETE("/profile", controllers.DeleteProfile) // Route to delete profile
	}

	r.Run()
}
