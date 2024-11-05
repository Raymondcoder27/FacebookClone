package initializers

import (
	"log"
	"os"

	"example.com/facebookclone/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	// dbHost := os.Getenv("DB_HOST")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)

	dsn := os.Getenv("DB")

	// log.Println("Connecting to database with DSN:", dsn)

	if dsn == "" {
		log.Fatal("Database url is not set")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
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
