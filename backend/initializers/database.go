package initializers

import (
	"log"

	"example.com/facebookclone/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.Comment{})
	if err != nil {
		log.Printf("Error migrating Comment Database: %v", err)
	}

	err2 := DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Printf("Error migrating Post Database: %v", err2)
	}

	err3 := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Printf("Error migrating User Database: %v", err3)
	}
}
