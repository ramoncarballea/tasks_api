package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"tasks.com/modules/base/dto"
	abstractServices "tasks.com/modules/task/domain/services"
	taskDto "tasks.com/modules/task/dto"
	"time"
)

// SuccessResponse represents a successful API response
// swagger:model SuccessResponse
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error API response
// swagger:model ErrorResponse
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// TaskDetailsResponse represents a task details response
// swagger:model TaskDetailsResponse
type TaskDetailsResponse struct {
	Success bool `json:"success"`
	Data    *struct {
		ID          uint       `json:"id"`
		Name        string     `json:"name"`
		Description string     `json:"description"`
		ExpiresAt   time.Time  `json:"expires_at"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   *time.Time `json:"updated_at"`
	} `json:"data"`
}

// CreateTaskRequest represents a task creation request
// swagger:model CreateTaskRequest
type CreateTaskRequest struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ExpiresAt   time.Time `json:"expires_at" validate:"required"`
}

// TaskListResponse represents a paginated list of tasks
// swagger:model TaskListResponse
type TaskListResponse struct {
	Success bool `json:"success"`
	Data    []*struct {
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ExpiresAt   time.Time `json:"expires_at"`
	} `json:"data"`
}

type TaskHandler struct {
	service abstractServices.TaskService
	log     *zap.Logger
}

func New(service abstractServices.TaskService, log *zap.Logger) *TaskHandler {
	return &TaskHandler{
		service: service,
		log:     log,
	}
}

// GetAll retrieves a paginated list of tasks
// @Summary Get all tasks
// @Description Get a paginated list of tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Param page_number query int true "Page number" minimum(1)
// @Param page_size query int true "Number of items per page" minimum(1) maximum(100)
// @Success 200 {object} TaskListResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/task [get]
func (h *TaskHandler) GetAll(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		c.JSON(400, dto.BadRequest("page_size is bad formatted"))
		return
	}

	pageNumber, err := strconv.Atoi(c.Query("page_number"))
	if err != nil {
		c.JSON(400, dto.BadRequest("page_number is bad formatted"))
		return
	}

	tasks, err := h.service.ListTasks(uint(pageNumber), uint(pageSize))
	if err != nil {
		h.log.Error("error listing tasks", zap.Error(err))
		c.JSON(500, dto.ServerError(fmt.Sprintf("error listing tasks: %v", err)))
		return
	}

	c.JSON(200, tasks)
}

// GetByID retrieves a task by its ID
// @Summary Get a task by ID
// @Description Get a single task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID" minimum(1)
// @Success 200 {object} TaskDetailsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/task/{id} [get]
func (h *TaskHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, dto.BadRequest("id is bad formatted"))
		return
	}

	task, err := h.service.GetTaskDetails(uint(id))
	if err != nil {
		h.log.Error("error getting task", zap.Error(err))
		c.JSON(500, dto.ServerError(fmt.Sprintf("error getting task: %v", err)))
		return
	}

	c.JSON(200, task)
}

// Create adds a new task
// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body CreateTaskRequest true "Task object that needs to be added"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/task [post]
func (h *TaskHandler) Create(c *gin.Context) {
	var model taskDto.CreateTaskDto
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(400, dto.BadRequest("invalid request"))
		return
	}

	if err := h.service.AddNewTask(model); err != nil {
		h.log.Error("error adding new task", zap.Error(err))
		c.JSON(500, dto.ServerError(fmt.Sprintf("error adding new task: %v", err)))
		return
	}

	c.JSON(200, gin.H{"message": "task created successfully"})
}

// Update updates an existing task
// @Summary Update a task
// @Description Update an existing task by ID with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID" minimum(1)
// @Param task body CreateTaskRequest true "Updated task object"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/task/{id} [put]
func (h *TaskHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, dto.BadRequest("id is bad formatted"))
		return
	}

	var model taskDto.CreateTaskDto
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(400, dto.BadRequest("invalid request"))
		return
	}

	if err := h.service.UpdateTask(uint(id), model); err != nil {
		h.log.Error("error updating task", zap.Error(err))
		c.JSON(500, dto.ServerError(fmt.Sprintf("error updating task: %v", err)))
		return
	}

	c.JSON(200, gin.H{"message": "task updated successfully"})
}

// Delete removes a task by ID
// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID" minimum(1)
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/task/{id} [delete]
func (h *TaskHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, dto.BadRequest("id is bad formatted"))
		return
	}

	if err := h.service.RemoveTask(uint(id)); err != nil {
		h.log.Error("error removing task", zap.Error(err))
		c.JSON(500, dto.ServerError(fmt.Sprintf("error removing task: %v", err)))
		return
	}

	c.JSON(200, gin.H{"message": "task removed successfully"})
}
