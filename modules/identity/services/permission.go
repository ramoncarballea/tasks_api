package services

import (
	domain "tasks.com/modules/base/domain/models"
	"tasks.com/modules/identity/domain/models"
	"tasks.com/modules/identity/domain/repositories"
	"tasks.com/modules/identity/domain/services"
	"tasks.com/modules/identity/dto"
	"tasks.com/utils/collections"
)

type permissionService struct {
	services.PermissionService
	repository repositories.PermissionRepository
}

func NewPermissionService(repository repositories.PermissionRepository) services.PermissionService {
	return &permissionService{
		repository: repository,
	}
}

func (p *permissionService) GetAll() ([]dto.PermissionDto, error) {
	permissions, err := p.repository.GetAll()
	if err != nil {
		return nil, err
	}

	dtos := collections.Map(permissions, func(permission *models.Permission) dto.PermissionDto {
		return dto.PermissionDto{
			ID:   permission.ID,
			Name: permission.Name,
		}
	})

	return dtos, nil
}

func (p *permissionService) GetDetails(name string) (*dto.PermissionDto, error) {
	permission, err := p.repository.GetByName(name)
	if err != nil {
		return nil, err
	}

	return &dto.PermissionDto{
		ID:   permission.ID,
		Name: permission.Name,
	}, nil
}

func (p *permissionService) Create(name string) (uint, error) {
	id, err := p.repository.Add(name)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *permissionService) Update(id uint, name string) error {
	err := p.repository.Update(models.Permission{
		BaseModel: domain.BaseModel[uint]{ID: id},
		Name:      name,
	})
	if err != nil {
		return err
	}

	return nil
}
