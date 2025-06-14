package repositories

import (
	"tasks.com/modules/identity/domain/models"
)

type RoleRepository interface {
	GetAll() ([]*models.Role, error)
	GetByName(name string) (*models.Role, error)
	Add(role models.Role) (uint, error)
	AddRange(roles []models.Role) error
	Update(role models.Role) error
	Delete(id uint) error
}
