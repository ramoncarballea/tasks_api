package services

import (
	domain "tasks.com/modules/base/domain/models"
	"tasks.com/modules/identity/domain/models"
	"tasks.com/modules/identity/domain/repositories"
	"tasks.com/modules/identity/domain/services"
	"tasks.com/modules/identity/dto"
	"tasks.com/utils/collections"
)

type roleService struct {
	services.RoleService
	repository repositories.RoleRepository
}

func NewRoleService(repository repositories.RoleRepository) services.RoleService {
	return &roleService{
		repository: repository,
	}
}

func (rs *roleService) Create(role dto.CreateRoleDto) (uint, error) {
	id, err := rs.repository.Add(models.Role{
		Name: role.Name,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (rs *roleService) Remove(id uint) error {
	err := rs.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (rs *roleService) Update(id uint, name string) error {
	err := rs.repository.Update(models.Role{
		BaseModel: domain.BaseModel[uint]{ID: id},
		Name:      name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (rs *roleService) GetAll() ([]dto.RoleDto, error) {
	roles, err := rs.repository.GetAll()
	if err != nil {
		return nil, err
	}

	dtos := collections.Map(roles, func(role *models.Role) dto.RoleDto {
		return dto.RoleDto{
			ID:   role.ID,
			Name: role.Name,
		}
	})

	return dtos, nil
}

func (rs *roleService) GetDetails(name string) (*dto.RoleDto, error) {
	role, err := rs.repository.GetByName(name)
	if err != nil {
		return nil, err
	}

	return &dto.RoleDto{
		ID:   role.ID,
		Name: role.Name,
	}, nil
}
