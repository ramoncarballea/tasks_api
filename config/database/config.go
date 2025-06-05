package database

import "go.uber.org/fx"

func ProvideDatabaseModule() fx.Option {
	return fx.Module(
		"database",
		fx.Provide(
			newPostgresqlDatabaseConnection,
		),
	)
}
