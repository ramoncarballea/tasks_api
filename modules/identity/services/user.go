package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"tasks.com/config/environment"
	domain "tasks.com/modules/base/domain/models"
	baseDto "tasks.com/modules/base/dto"
	"tasks.com/modules/identity/domain/models"
	"tasks.com/modules/identity/domain/repositories"
	"tasks.com/modules/identity/domain/services"
	"tasks.com/modules/identity/dto"
	"tasks.com/utils/collections"
)

type userService struct {
	services.UserService
	config     *environment.PasswordConfig
	repository repositories.UserRepository
}

func NewUserService(r repositories.UserRepository, config *environment.PasswordConfig) services.UserService {
	return &userService{
		repository: r,
		config:     config,
	}
}

func (us *userService) GetAll(pageNumber uint, pageSize uint) (*baseDto.PagedResponse[dto.ListUserDto], error) {
	users, total, err := us.repository.GetAll(pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}

	dtos := collections.Map(users, func(user *models.User) dto.ListUserDto {
		return dto.ListUserDto{
			ID:       user.ID,
			Email:    user.Email,
			FullName: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		}
	})

	return &baseDto.PagedResponse[dto.ListUserDto]{
		Data:       dtos,
		PageNumber: pageNumber,
		PageSize:   pageSize,
		TotalItems: int64(total),
	}, nil
}

func (us *userService) GetByID(id uuid.UUID) (*dto.UserDetailsDto, error) {
	user, err := us.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	permissions := make([]dto.PermissionDto, 0)
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			permissions = append(permissions, dto.PermissionDto{
				ID:   permission.ID,
				Name: permission.Name,
			})
		}
	}

	roles := collections.Map(user.Roles, func(role *models.Role) dto.RoleDto {
		return dto.RoleDto{
			ID:   role.ID,
			Name: role.Name,
		}
	})

	return &dto.UserDetailsDto{
		ID:          user.ID,
		Email:       user.Email,
		FullName:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Permissions: permissions,
		Roles:       roles,
	}, nil
}

func (us *userService) Create(model dto.CreateUserDto) (uuid.UUID, error) {
	passwordHash := us.hashPassword(model.Password)

	user := models.User{
		Roles: collections.Map(model.Roles, func(role dto.RoleDto) *models.Role {
			return &models.Role{
				BaseModel: domain.BaseModel[uint]{ID: role.ID},
				Name:      role.Name,
			}
		}),
		LastName:  model.LastName,
		FirstName: model.FirstName,
		Email:     model.Email,
		Password:  passwordHash,
	}

	id, err := us.repository.Add(user)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (us *userService) Update(id uuid.UUID, model dto.CreateUserDto) error {
	err := us.repository.Update(models.User{
		Roles: collections.Map(model.Roles, func(role dto.RoleDto) *models.Role {
			return &models.Role{
				BaseModel: domain.BaseModel[uint]{ID: role.ID},
				Name:      role.Name,
			}
		}),
		LastName:  model.LastName,
		FirstName: model.FirstName,
		Email:     model.Email,
		Password:  model.Password,
		BaseModel: domain.BaseModel[uuid.UUID]{ID: id},
	})

	if err != nil {
		return err
	}

	return nil
}

func (us *userService) Remove(id uuid.UUID) error {
	err := us.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (us *userService) hashPassword(password string) string {
	key := us.config.SecretKey
	passwordHash := sha256.Sum256([]byte(fmt.Sprintf("%s.%s", password, key)))
	result := hex.EncodeToString(passwordHash[:])
	return result
}
