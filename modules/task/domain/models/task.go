package models

import (
	"time"

	baseModels "tasks.com/modules/base/domain/models"
)

type Task struct {
	baseModels.BaseModel[uint]
	baseModels.Auditable
	Name        string
	Description string
	ExpiresAt   time.Time
}
