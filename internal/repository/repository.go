package repository

import (
	"github.com/JuanPO17/myfirstgo/internal/model"
	"gorm.io/gorm"
)

// TaskRepository defines the interface for interacting with task data.
type TaskRepository interface {
	CreateTask(task *model.Task) error
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(id uint) (*model.Task, error)
	UpdateTask(task *model.Task) error
	DeleteTask(id uint) error
}

// taskRepository is the implementation of TaskRepository.
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new TaskRepository.
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// CreateTask creates a new task in the database.
func (r *taskRepository) CreateTask(task *model.Task) error {
	return r.db.Create(task).Error
}

// GetAllTasks retrieves all tasks from the database.
func (r *taskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

// GetTaskByID retrieves a single task from the database by its ID.
func (r *taskRepository) GetTaskByID(id uint) (*model.Task, error) {
	var task model.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// UpdateTask updates an existing task in the database.
func (r *taskRepository) UpdateTask(task *model.Task) error {
	return r.db.Save(task).Error
}

// DeleteTask removes a task from the database by its ID.
func (r *taskRepository) DeleteTask(id uint) error {
	return r.db.Delete(&model.Task{}, id).Error
}
