package repositories

import (
	"github.com/google/uuid"
	"tasks.com/modules/identity/domain/models"
)

type PermissionRepository interface {
	GetAll() ([]*models.Permission, error)
	GetByUserID(id uuid.UUID) ([]*models.Permission, error)
	GetByName(name string) (*models.Permission, error)
	Add(name string) (uint, error)
	AddRange(permissions []models.Permission) error
	Update(permission models.Permission) error
	Delete(id uint) error
}
