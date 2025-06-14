package identity

import (
	"go.uber.org/fx"
	"tasks.com/modules/identity/handlers"
	"tasks.com/modules/identity/repositories"
	"tasks.com/modules/identity/routes"
	"tasks.com/modules/identity/services"
)

func ProvideIdentityModule() fx.Option {
	return fx.Module(
		"identity",
		fx.Provide(
			repositories.NewRoleRepository,
			repositories.NewUserRepository,
			repositories.NewPermissionRepository,
			services.NewRoleService,
			services.NewPermissionService,
			services.NewUserService,
			handlers.NewUserHandler,
		),
		routes.ProvideUserRoutes(),
	)
}
