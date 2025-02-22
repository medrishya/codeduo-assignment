// handlers/todo.go
package handlers

import (
	"codeduo-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get all ToDos
func GetTasks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the pagination parameters from query string
        pageStr := c.Query("page")
        limitStr := c.Query("limit")

        // Set default values if not provided
        page := 1
        limit := 10

		if pageStr != "" {
            parsedPage, err := strconv.Atoi(pageStr)
            if err == nil && parsedPage > 0 {
                page = parsedPage
            }
        }
        if limitStr != "" {
            parsedLimit, err := strconv.Atoi(limitStr)
            if err == nil && parsedLimit > 0 {
                limit = parsedLimit
            }
        }
		
		var tasks []models.Task
		var total int64

		// Fetch the total count of tasks
        db.Model(&models.Task{}).Count(&total)

        // Fetch todos with pagination
        offset := (page - 1) * limit
        if err := db.Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve tasks."})
            return
        }

        // Return the paginated response
        c.JSON(http.StatusOK, gin.H{
            "total": total,
            "page":  page,
            "limit": limit,
            "todos": tasks,
        })
	}
}

// Get a ToDo by ID
func GetTaskByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var task models.Task

		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}

// Create a new ToDo
func CreateTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(task)
		db.Create(&task)
		// Start a goroutine to handle post-creation actions
        go func(newTask models.Task) {
			// For now just logging in the print details on the console.
            // Here you could add logic to send notifications, etc.
			log.Info().
			Interface("task", newTask). // Use Interface to log the whole struct
			Msg("Created new task")

        }(task) 
		c.JSON(http.StatusCreated, task)
	}
}

// Update an existing ToDo
func UpdateTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		id := c.Param("id")
		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&task)
		c.JSON(http.StatusOK, task)
	}
}

// Delete a ToDo
func DeleteTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Task{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
