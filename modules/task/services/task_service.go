package services

import (
	"fmt"

	"go.uber.org/zap"
	domain "tasks.com/modules/base/domain/models"
	baseDto "tasks.com/modules/base/dto"
	"tasks.com/modules/task/domain/models"
	abstractRepositories "tasks.com/modules/task/domain/repositories"
	abstractServices "tasks.com/modules/task/domain/services"
	"tasks.com/modules/task/dto"
)

type taskService struct {
	abstractServices.TaskService
	r   abstractRepositories.TaskRepository
	log *zap.Logger
}

func NewTaskService(r abstractRepositories.TaskRepository, log *zap.Logger) abstractServices.TaskService {
	return &taskService{r: r, log: log}
}

func (s *taskService) ListTasks(pageNumber, pageSize uint) (*baseDto.PagedResponse[dto.ListTaskDto], error) {
	s.log.Info("listing tasks with", zap.Uint("page_number", pageNumber), zap.Uint("page_size", pageSize))
	tasks, total, err := s.r.GetAll(pageSize, pageSize*(pageNumber-1))
	if err != nil {
		s.log.Error("error while getting tasks", zap.Error(err))
		return nil, err
	}

	s.log.Info(fmt.Sprintf("found %d tasks", total))

	taskDtos := make([]dto.ListTaskDto, 0)
	for _, task := range tasks {
		taskDtos = append(taskDtos, dto.ListTaskDto{
			ID:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			ExpiresAt:   task.ExpiresAt,
		})
	}

	response := baseDto.PagedResponse[dto.ListTaskDto]{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		TotalItems: int64(total),
		Data:       taskDtos,
	}

	s.log.Info("list tasks response", zap.Any("response", response))

	return &response, nil
}

func (s *taskService) AddNewTask(task dto.CreateTaskDto) error {
	s.log.Info("trying to add new task", zap.Any("task", task))
	if _, err := s.r.Add(models.Task{
		Name:        task.Name,
		ExpiresAt:   task.ExpiresAt,
		Description: task.Description,
	}); err != nil {
		s.log.Error("error while adding new task", zap.Error(err))
		return err
	}

	return nil
}

func (s *taskService) RemoveTask(id uint) error {
	s.log.Info("trying to remove task with id", zap.Uint("id", id))
	if err := s.r.Delete(id); err != nil {
		s.log.Error("error adding new task", zap.Error(err))
		return err
	}
	return nil
}

func (s *taskService) UpdateTask(id uint, task dto.CreateTaskDto) error {
	s.log.Info("trying to update task with id", zap.Uint("id", id), zap.Any("task", task))
	if err := s.r.Update(models.Task{
		BaseModel:   domain.BaseModel[uint]{ID: id},
		Name:        task.Name,
		Description: task.Description,
		ExpiresAt:   task.ExpiresAt,
	}); err != nil {
		s.log.Error("error while updating task", zap.Error(err))
		return err
	}

	return nil
}

func (s *taskService) GetTaskDetails(id uint) (*dto.TaskDetailsDto, error) {
	s.log.Info("trying to get task details by id", zap.Uint("id", id))
	task, err := s.r.GetByID(id)
	s.log.Info("task obtained", zap.Any("task", task))
	if err != nil {
		s.log.Error("error while getting task by id", zap.Error(err))
		return nil, err
	}

	return &dto.TaskDetailsDto{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		ExpiresAt:   task.ExpiresAt,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}
