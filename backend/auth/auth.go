package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"example.com/facebookclone/initializers"
	"example.com/facebookclone/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}
	c.BindJSON(&body)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to hash password."})
		return
	}

	placeholderImagePath := "public/user-placeholder.png"

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash), Image: placeholderImagePath}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create user."})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}

// func Login(c *gin.Context) {
// 	//Get the email and password off the request body
// 	var body struct {
// 		Email    string
// 		Password string
// 	}

// 	//Bind the request body to the body struct
// 	if c.Bind(&body) != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password."})
// 		return
// 	}

// 	//Look up requested user
// 	var user models.User
// 	initializers.DB.First(&user, "email = ?", body.Email)

// 	if user.ID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found."})
// 		return
// 	}

// 	//compare passed in password with saved password
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to hash password"})
// 		return
// 	}

// 	//Generate a JWT token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": user.ID,
// 		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
// 	})

// 	//sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create token."})
// 		return
// 	}

// 	//return with the user
// 	c.JSON(http.StatusOK, gin.H{"token": tokenString})
// }

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	// if c.BindJSON(&body) != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password."})
	// 	return
	// }

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password.", "error": err.Error()})
		return
	}

	// Look up user by email
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found."})
		return
	}

	// Compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect password."})
		return
	}

	fmt.Println("Stored password hash:", user.Password)
	fmt.Println("Incoming password:", body.Password)

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create token."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Validate(c *gin.Context) {
	// fmt.Println("I'm logged in")
	user, _ := c.Get("user")

	// c.JSON(http.StatusOK, gin.H{
	// 	"user": user,
	// })
	c.JSON(http.StatusOK, user)
}
