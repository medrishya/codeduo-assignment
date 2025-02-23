package models

import (
	"time"

	"gorm.io/gorm"
)

// TaskStatus represents the status of a Task item
type TaskStatus string

const (
	Pending  TaskStatus = "pending"
	Completed     TaskStatus = "completed"
)

// Task model
type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" binding:"required"`
	Status      TaskStatus `json:"status" gorm:"type:text;default:'pending'" binding:"required,oneof=pending completed"`
	CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}

// Migrate the schema
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}