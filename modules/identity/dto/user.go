package dto

import (
	_ "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type ListUserDto struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
}

type UserDetailsDto struct {
	ID          uuid.UUID       `json:"id"`
	FullName    string          `json:"full_name"`
	Email       string          `json:"email"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at"`
	Permissions []PermissionDto `json:"permissions"`
	Roles       []RoleDto       `json:"roles"`
}

type CreateUserDto struct {
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Roles     []RoleDto `json:"roles" validate:"required"`
}

type SignUpDto struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type SignInDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
