package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JuanPO17/myfirstgo/internal/model"
	"github.com/JuanPO17/myfirstgo/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTest(t *testing.T) (*gin.Engine, repository.TaskRepository) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Task{})
	assert.NoError(t, err)

	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})

	repo := repository.NewTaskRepository(db)
	handler := NewTaskHandler(repo)
	router := gin.Default()
	router.POST("/tasks", handler.CreateTask)
	router.GET("/tasks", handler.GetAllTasks)
	router.GET("/tasks/:id", handler.GetTaskByID)
	router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)

	return router, repo
}

func TestCreateTask(t *testing.T) {
	router, _ := setupTest(t)

	createReq := model.CreateTaskRequest{Name: "New Task", Description: "New Description"}
	body, _ := json.Marshal(createReq)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var task model.Task
	err := json.Unmarshal(w.Body.Bytes(), &task)
	assert.NoError(t, err)
	assert.Equal(t, "New Task", task.Name)
}

func TestGetAllTasks(t *testing.T) {
	router, repo := setupTest(t)

	repo.CreateTask(&model.Task{Name: "Task 1"})

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var tasks []model.Task
	err := json.Unmarshal(w.Body.Bytes(), &tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
}

func TestGetTaskByID(t *testing.T) {
	router, repo := setupTest(t)

	task := &model.Task{Name: "Task 1"}
	repo.CreateTask(task)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var foundTask model.Task
	err := json.Unmarshal(w.Body.Bytes(), &foundTask)
	assert.NoError(t, err)
	assert.Equal(t, "Task 1", foundTask.Name)
}

func TestUpdateTask(t *testing.T) {
	router, repo := setupTest(t)

	task := &model.Task{Name: "Original Name"}
	repo.CreateTask(task)

	updateReq := model.UpdateTaskRequest{Name: new(string)}
	*updateReq.Name = "Updated Name"
	body, _ := json.Marshal(updateReq)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedTask model.Task
	err := json.Unmarshal(w.Body.Bytes(), &updatedTask)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedTask.Name)
}

func TestDeleteTask(t *testing.T) {
	router, repo := setupTest(t)

	task := &model.Task{Name: "To Be Deleted"}
	repo.CreateTask(task)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
