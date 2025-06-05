package dto

import "time"

type ListTaskDto struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expires_at"`
}
