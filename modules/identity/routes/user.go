package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"tasks.com/modules/identity/handlers"
)

func ProvideUserRoutes() fx.Option {
	return fx.Invoke(
		func(server *gin.Engine, userHandler *handlers.UserHandler) {
			group := server.Group("api/v1/user")
			{
				group.POST("/signup", userHandler.SignUp)
			}
		},
	)
}
