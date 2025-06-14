package domain

type BaseModel[T any] struct {
	ID T `db:"id"`
}
