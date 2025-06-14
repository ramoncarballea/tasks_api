package services

import (
	"tasks.com/modules/identity/dto"
)

type PermissionService interface {
	GetAll() ([]dto.PermissionDto, error)
	GetDetails(name string) (*dto.PermissionDto, error)
	Create(name string) (uint, error)
	Update(id uint, name string) error
	Remove(id uint) error
}
