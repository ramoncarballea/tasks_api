package app

import (
	"context"
	"errors"
	"fmt"
	"tasks.com/config/cache"
	"tasks.com/modules/base/routes"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"tasks.com/config/database"
	"tasks.com/config/environment"
	"tasks.com/modules/task"
)

func BuildApp() *fx.App {
	return fx.New(
		environment.ProvideEnvironment(),
		database.ProvideDatabaseModule(),
		fx.Provide(
			func() *gin.Engine {
				return gin.Default()
			},
			cache.NewMemory,
			zap.NewExample,
		),
		task.ProvideTaskModule(),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(autoMigrate),
		fx.Invoke(routes.ProvideBaseHandlers),
		fx.Invoke(newHttpServer),
	)
}

func newHttpServer(lc fx.Lifecycle, handler *gin.Engine, log *zap.Logger, config *environment.ServerConfig) *http.Server {
	addr := fmt.Sprintf(":%s", config.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func(server *http.Server) {
				err := server.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error("error starting http server", zap.Error(err))
					panic(err)
				}
			}(server)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("shutting down http server")
			return server.Shutdown(ctx)
		},
	})

	return server
}

func autoMigrate(con database.Connection, config *environment.DataBaseConfig, log *zap.Logger) {
	if !config.AutoMigrate {
		return
	}

	log.Info("applying migrations")
	if err := con.ApplyMigrations(); err != nil {
		panic(fmt.Sprintf("error while applying migrations: %s", err.Error()))
	}
}
