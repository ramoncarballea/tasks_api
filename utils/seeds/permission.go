package seeds

import (
	"go.uber.org/zap"
	"tasks.com/modules/identity/domain/models"
	"tasks.com/modules/identity/domain/repositories"
	"tasks.com/utils/collections"
)

type PermissionSeeder struct {
	Seeder
	repository repositories.PermissionRepository
	log        *zap.Logger
}

func NewPermissionSeeder(repository repositories.PermissionRepository, log *zap.Logger) Seeder {
	return &PermissionSeeder{
		repository: repository,
		log:        log,
	}
}

func (ps *PermissionSeeder) Seed() {
	permissions, err := ps.repository.GetAll()
	if err != nil || len(permissions) != 0 {
		ps.log.Info("permissions already exist")
		return
	}

	ps.log.Info("seeding permissions", zap.Strings("user", UserPermissions), zap.Strings("role", RolePermissions), zap.Strings("permission", PermissionPermissions), zap.Strings("task", TaskPermissions))

	allPermissions := append(UserPermissions, RolePermissions...)
	allPermissions = append(allPermissions, PermissionPermissions...)
	allPermissions = append(allPermissions, TaskPermissions...)

	baseModels := collections.Map(allPermissions, func(name string) models.Permission {
		return models.Permission{
			Name: name,
		}
	})

	if err := ps.repository.AddRange(baseModels); err != nil {
		ps.log.Error("failed to add permissions", zap.Error(err))
		panic(err)
	}
}
