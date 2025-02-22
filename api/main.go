package main

import (
	"codeduo-api/database"
	"codeduo-api/models"
	"codeduo-api/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Set up ZeroLog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	r := gin.Default()

	// Connect to database
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Task{})

	// Setup routes
	routes.SetupRoutes(r, database.DB)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, from api application of CodeDuo"})
	})

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}
	
}