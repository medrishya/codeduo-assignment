package cli

import (
	"codeduo-api/models"

	"github.com/rs/zerolog/log"

	"os"

	"codeduo-api/routes"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

// Root command
var rootCmd = &cobra.Command{
    Use:   "Task",
    Short: "A simple task application",
}

// Add command
var addCmd = &cobra.Command{
    Use:   "add [name]",
    Short: "Add a new task",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        status := models.Pending

        // Create new Todo
        task := models.Task{
            Name:        name,
            Status:      status,
        }

        if err := db.Create(&task).Error; err != nil {
			log.Fatal().Err(err).Msg("Error adding task : ")
        }

		log.Info().
		Interface("task", task). // Use Interface to log the whole struct
		Msg("New task added:  ")
		os.Exit(0);
    },
}

// List command
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        var tasks []models.Task

        // Fetch all todos
        if err := db.Find(&tasks).Error; err != nil {
			log.Fatal().Err(err).Msg("Error retrieving task : ")
        }

        // Display todos
        for _, task := range tasks {
			log.Info().
			Interface("task", task). 
			Msg("Task Details :  ")
        }
		os.Exit(0);
    },
}

// Process command
var processCmd = &cobra.Command{
    Use:   "process",
    Short: "Start a worker to process a task from the database and exit once completed",
    Run: func(cmd *cobra.Command, args []string) {
        go worker()
    },
}

// Worker function to process a task
func worker() {
    var task models.Task

    // Fetch the first task pending
    statuses := []string{string(models.Pending)}
    if err := db.Where("status IN ?", statuses).First(&task).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
			log.Info().Msg("No task to process. Exiting ... ")
            return
        } else {
			log.Fatal().Err(err).Msg("Error retrieving task : ")
        }
    }

    // Mark the task as done
    task.Status = models.Completed

    // Update the task in the database
    if err := db.Save(&task).Error; err != nil {
		log.Fatal().Err(err).Msg("Error updating task : ")
    } else {
		log.Info().
		Interface("task", task). 
		Msg("Processed task details :  ")
    }

	os.Exit(0)
}

// To start the rest api
var apiCmd = &cobra.Command{
    Use:   "api",
    Short: "Run the API server",
    Run: func(cmd *cobra.Command, args []string) {
        // Initialize Gin router
        router := gin.Default()

        // Set up routes
        routes.SetupRoutes(router, db)

        // Start the server
        if err := router.Run(":8080"); err != nil {
			log.Fatal().Err(err).Msg("Error starting the API server: ")
        }
    },
}

// Initialize the CLI commands
func Init() {
    var err error

    // Initialize the database
    db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database :")
    }

    // Run migrations
    if err := models.Migrate(db); err != nil {
		log.Fatal().Err(err).Msg("Migration Failed :")
    }

    rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(processCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(apiCmd)
}

// Execute runs the root command
func Execute() {
    if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Error Executing command :")
        os.Exit(1)
    }
}
