package seeds

import "go.uber.org/fx"

func ProvideSeedsModule() fx.Option {
	return fx.Module(
		"seeds",
		fx.Provide(
			fx.Annotate(NewPermissionSeeder, fx.As(new(Seeder)), fx.ResultTags(`name:"permissionSeeder"`)),
			fx.Annotate(NewRoleSeeder, fx.As(new(Seeder)), fx.ResultTags(`name:"roleSeeder"`)),
			fx.Annotate(NewUserSeeder, fx.As(new(Seeder)), fx.ResultTags(`name:"userSeeder"`)),
		),
		fx.Invoke(
			fx.Annotate(
				func(permissionSeeder Seeder, roleSeeder Seeder, userSeeder Seeder) {
					permissionSeeder.Seed()
					roleSeeder.Seed()
					userSeeder.Seed()
				}, fx.ParamTags(`name:"permissionSeeder"`, `name:"roleSeeder"`, `name:"userSeeder"`)),
		),
	)
}
