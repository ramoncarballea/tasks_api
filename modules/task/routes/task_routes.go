package routes

import (
	"github.com/gin-gonic/gin"
	taskHandler "tasks.com/modules/task/handlers"
)

func ProvideRoutes(router *gin.Engine, handler *taskHandler.TaskHandler) {
	group := router.Group("api/v1/task")
	{
		group.GET("/", handler.GetAll)
		group.GET("/:id", handler.GetByID)
		group.POST("/", handler.Create)
		group.PUT("/:id", handler.Update)
		group.DELETE("/:id", handler.Delete)
	}
}
