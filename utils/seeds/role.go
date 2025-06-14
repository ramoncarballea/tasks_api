package seeds

import (
	"go.uber.org/zap"
	"tasks.com/modules/identity/domain/models"
	"tasks.com/modules/identity/domain/repositories"
	"tasks.com/utils/collections"
)

type RoleSeeder struct {
	Seeder
	repository  repositories.RoleRepository
	pRepository repositories.PermissionRepository
	log         *zap.Logger
}

func NewRoleSeeder(log *zap.Logger, pRepository repositories.PermissionRepository, repository repositories.RoleRepository) Seeder {
	return &RoleSeeder{
		repository:  repository,
		log:         log,
		pRepository: pRepository,
	}
}

func (rs *RoleSeeder) Seed() {
	existingRoles, err := rs.repository.GetAll()
	if err != nil || len(existingRoles) != 0 {
		rs.log.Info("roles already exist")
		return
	}

	allPermissions, err := rs.pRepository.GetAll()
	if err != nil {
		panic(err)
	}

	parsePermission := func(name string) *models.Permission {
		for _, p := range allPermissions {
			if p.Name == name {
				return p
			}
		}

		return nil
	}

	admin := append(append(append(UserPermissions, RolePermissions...), TaskPermissions...), PermissionPermissions...)

	userPermissions := collections.Map(append([]string{"USER_READ", "USER_UPDATE"}, TaskPermissions...), parsePermission)
	adminPermissions := collections.Map(admin, parsePermission)
	var roles = []models.Role{
		{
			Name:        "admin",
			Permissions: adminPermissions,
		},
		{
			Name:        "user",
			Permissions: userPermissions,
		},
	}

	if err := rs.repository.AddRange(roles); err != nil {
		rs.log.Info("failed to add roles", zap.Error(err))
		panic(err)
	}
}
