package database

import (
	"codeduo-api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected successfully!")
}

func MigrateDatabase() {
	DB.AutoMigrate(&models.Task{})
	log.Println("Database migrated successfully!")
}