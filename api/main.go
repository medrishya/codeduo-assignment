package main

import (
	"codeduo-api/database"
	"codeduo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Task{})
	// Routes
	r.GET("/tasks", getTasks)
	r.POST("/tasks", createTask)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, from Gin!"})
	})

	r.Run(":8080") // Start server on port 8080
}

// Get all users
func getTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// Create a new user
func createTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}