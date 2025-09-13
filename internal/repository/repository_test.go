package repository

import (
	"testing"

	"github.com/JuanPO17/myfirstgo/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Task{})
	assert.NoError(t, err)

	return db
}

func TestTaskRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	// Test CreateTask
	task := &model.Task{Name: "Test Task", Description: "Test Description"}
	err := repo.CreateTask(task)
	assert.NoError(t, err)
	assert.NotZero(t, task.ID)

	// Test GetTaskByID
	foundTask, err := repo.GetTaskByID(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Test Task", foundTask.Name)

	// Test GetAllTasks
	tasks, err := repo.GetAllTasks()
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)

	// Test UpdateTask
	foundTask.Name = "Updated Task"
	foundTask.Completed = true
	err = repo.UpdateTask(foundTask)
	assert.NoError(t, err)

	updatedTask, err := repo.GetTaskByID(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", updatedTask.Name)
	assert.True(t, updatedTask.Completed)

	// Test DeleteTask
	err = repo.DeleteTask(task.ID)
	assert.NoError(t, err)

	_, err = repo.GetTaskByID(task.ID)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
