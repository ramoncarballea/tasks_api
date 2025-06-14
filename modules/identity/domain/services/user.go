package services

import (
	"github.com/google/uuid"
	baseDto "tasks.com/modules/base/dto"
	"tasks.com/modules/identity/dto"
)

type UserService interface {
	GetAll(pageNumber uint, pageSize uint) (*baseDto.PagedResponse[dto.ListUserDto], error)
	GetDetails(id uuid.UUID) (dto.UserDetailsDto, error)
	Create(model dto.CreateUserDto) (uuid.UUID, error)
	Update(id uuid.UUID, model dto.CreateUserDto) error
	Remove(id uuid.UUID) error
}
