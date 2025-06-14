package dto

type RoleDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateRoleDto struct {
	Name        string   `json:"name" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}
