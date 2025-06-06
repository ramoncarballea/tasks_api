package swagger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tasks.com/config/environment"
	"tasks.com/docs"
)

func SetupSwagger(handler *gin.Engine, config *environment.ServerConfig) {
	docs.SwaggerInfo.Title = "Tasks API"
	docs.SwaggerInfo.Description = "This is a simple task management API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.Host, config.Port)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
