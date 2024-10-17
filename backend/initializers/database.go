package initializers

import "gorm.io/gorm"

var DB *gorm.DB

func ConnectToDB() {
	var err error
}

func MigrateDB() {
	err := DB.AutoMigrate()
}
