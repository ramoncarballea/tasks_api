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
	"tasks.com/config/swagger"
	_ "tasks.com/docs" // Import docs for Swagger
	"tasks.com/modules/task"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
		fx.Invoke(swagger.SetupSwagger),
		fx.Invoke(routes.ProvideBaseHandlers),
		fx.Invoke(newHttpServer),
	)
}

func newHttpServer(lc fx.Lifecycle, handler *gin.Engine, log *zap.Logger, config *environment.ServerConfig) *http.Server {
	// Setup Swagger
	log.Info("Swagger documentation available at /swagger/index.html")
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
