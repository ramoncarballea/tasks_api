package repositories

import (
	"github.com/google/uuid"
	"tasks.com/modules/identity/domain/models"
)

type UserRepository interface {
	GetAll(take, offset uint) ([]*models.User, uint, error)
	GetByID(id uuid.UUID) (*models.User, error)
	Add(user models.User) (uuid.UUID, error)
	Update(user models.User) error
	Delete(id uuid.UUID) error
}
