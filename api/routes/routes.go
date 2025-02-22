// routes/routes.go
package routes

import (
	"codeduo-api/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures the routes for the application
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/tasks", handlers.GetTasks(db))
	r.GET("/tasks/:id", handlers.GetTaskByID(db))
	r.POST("/tasks", handlers.CreateTask(db))
	r.PUT("/tasks/:id", handlers.UpdateTask(db))
	r.DELETE("/tasks/:id", handlers.DeleteTask(db))
}
