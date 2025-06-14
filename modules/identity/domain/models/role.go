package models

import (
	"tasks.com/modules/base/domain/models"
)

type Role struct {
	domain.BaseModel[uint]
	Name        string `db:"name"`
	Permissions []*Permission
}
