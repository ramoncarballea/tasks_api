package services

import (
	"tasks.com/modules/identity/dto"
)

type RoleService interface {
	GetAll() ([]dto.RoleDto, error)
	GetDetails(name string) (*dto.RoleDto, error)
	Create(role dto.CreateRoleDto) (uint, error)
	Update(id uint, name string) error
	Remove(id uint) error
}
