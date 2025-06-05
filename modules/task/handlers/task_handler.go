package handler

import (
	"fmt"
	"go.uber.org/zap"
	"strconv"

	"github.com/gin-gonic/gin"
	"tasks.com/modules/base/dto"
	abstractServices "tasks.com/modules/task/domain/services"
	taskDto "tasks.com/modules/task/dto"
)

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
