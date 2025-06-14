package models

import (
	"github.com/google/uuid"
	domain "tasks.com/modules/base/domain/models"
)

type User struct {
	domain.BaseModel[uuid.UUID]
	domain.Auditable
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Roles     []*Role
}
