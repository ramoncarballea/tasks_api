package domain

import "time"

type Auditable struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
}
