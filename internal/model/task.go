package model

import "time"

// Task represents a task in the system.
type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null" example:"Buy milk"`
	Description string    `json:"description" example:"Get a gallon of whole milk"`
	Completed   bool      `json:"completed" gorm:"default:false" example:"false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateTaskRequest defines the shape of a request to create a new task.
type CreateTaskRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// UpdateTaskRequest defines the shape of a request to update an existing task.
type UpdateTaskRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}
