package task

import (
	"go.uber.org/fx"
	"tasks.com/modules/task/handlers"
	"tasks.com/modules/task/repositories"
	"tasks.com/modules/task/routes"
	"tasks.com/modules/task/services"
)

func ProvideTaskModule() fx.Option {
	return fx.Module(
		"task",
		fx.Provide(
			repositories.NewTaskRepository,
			services.NewTaskService,
			handler.New,
		),
		fx.Invoke(
			routes.ProvideRoutes,
		),
	)
}
