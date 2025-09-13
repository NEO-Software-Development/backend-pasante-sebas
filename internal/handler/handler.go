package handler

import (
	"net/http"
	"strconv"

	"github.com/JuanPO17/myfirstgo/internal/model"
	"github.com/JuanPO17/myfirstgo/internal/repository"
	"github.com/gin-gonic/gin"
)

// TaskHandler holds the repository for task operations.
type TaskHandler struct {
	repo repository.TaskRepository
}

// NewTaskHandler creates a new TaskHandler.
func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   task body model.CreateTaskRequest true "Create Task"
// @Success 201 {object} model.Task
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req model.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	task := &model.Task{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.repo.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Get a list of all tasks
// @Tags tasks
// @Produce  json
// @Success 200 {array} model.Task
// @Failure 500 {object} model.ErrorResponse
// @Router /tasks [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.repo.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID godoc
// @Summary Get a task by ID
// @Description Get a single task by its ID
// @Tags tasks
// @Produce  json
// @Param   id   path      int  true  "Task ID"
// @Success 200 {object} model.Task
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := h.repo.GetTaskByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update an existing task with the input payload
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   id   path      int  true  "Task ID"
// @Param   task body      model.UpdateTaskRequest true "Update Task"
// @Success 200 {object} model.Task
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	var req model.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	task, err := h.repo.GetTaskByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Task not found"})
		return
	}

	if req.Name != nil {
		task.Name = *req.Name
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Completed != nil {
		task.Completed = *req.Completed
	}

	if err := h.repo.UpdateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags tasks
// @Param   id   path      int  true  "Task ID"
// @Success 204
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	if err := h.repo.DeleteTask(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to delete task"})
		return
	}

	c.Status(http.StatusNoContent)
}
