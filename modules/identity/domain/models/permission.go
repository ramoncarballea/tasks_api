package models

import (
	domain "tasks.com/modules/base/domain/models"
)

type Permission struct {
	domain.BaseModel[uint]
	Name string `db:"name"`
}
