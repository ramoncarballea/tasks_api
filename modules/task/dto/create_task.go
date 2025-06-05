package dto

import "time"

type CreateTaskDto struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ExpiresAt   time.Time `json:"expires_at" validate:"required"`
}
