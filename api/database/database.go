package database

import (
	"codeduo-api/models"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database:")
	}
	log.Info().Msg("Database connected successfully!")
}

func MigrateDatabase() {
	DB.AutoMigrate(&models.Task{})
	log.Info().Msg("Database migrated successfully!")
}