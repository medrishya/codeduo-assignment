package main

import (
	"codeduo-api/cli"
	"codeduo-api/database"
	"codeduo-api/models"
	"codeduo-api/routes"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize the CLI commands and database
	cli.Init()       
    cli.Execute() 

	// CORS configuration to allow all origins
    corsConfig := cors.Config{
        // AllowAllOrigins: true, // Allows all origins
		AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // Allowed headers
        ExposeHeaders:    []string{"Content-Length"}, // Expose these headers to the browser
        MaxAge:          12 * 3600, // Cache duration in seconds
    }

	// Set up ZeroLog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// GIN Framework
	r := gin.Default()
	
	// Connect to database
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Task{})
	
	// Setup routes
	r.Use(cors.New(corsConfig))
	routes.SetupRoutes(r, database.DB)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, from api application of CodeDuo"})
	})

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}
	
}