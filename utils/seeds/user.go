package seeds

import (
	"fmt"
	"go.uber.org/zap"
	"tasks.com/config/environment"
	"tasks.com/modules/identity/domain/services"
	"tasks.com/modules/identity/dto"
)

type UserSeeder struct {
	Seeder
	log     *zap.Logger
	service services.UserService
	config  *environment.AdminUserConfig
}

func NewUserSeeder(log *zap.Logger, service services.UserService) Seeder {
	return &UserSeeder{
		log:     log,
		service: service,
		config:  environment.NewAdminUserConfig(),
	}
}

func (us *UserSeeder) Seed() {
	users, err := us.service.GetAll(1, 1)
	if err != nil || users.TotalItems != 0 {
		us.log.Info("users already exist")
		return
	}

	adminUser := dto.CreateUserDto{
		FirstName: "admin",
		LastName:  "admin",
		Email:     us.config.Email,
		Password:  us.config.Password,
		Roles: []dto.RoleDto{
			{
				ID:   1,
				Name: "admin",
			},
		},
	}

	_, err = us.service.Create(adminUser)
	if err != nil {
		panic(fmt.Sprintf("error while creating admin user: %s", err.Error()))
	}
}
