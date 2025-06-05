package services

import (
	baseDto "tasks.com/modules/base/dto"
	"tasks.com/modules/task/dto"
)

type TaskService interface {
	ListTasks(pageNumber uint, pageSize uint) (*baseDto.PagedResponse[dto.ListTaskDto], error)
	AddNewTask(task dto.CreateTaskDto) error
	RemoveTask(id uint) error
	UpdateTask(id uint, task dto.CreateTaskDto) error
	GetTaskDetails(id uint) (*dto.TaskDetailsDto, error)
}
