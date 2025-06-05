package routes

import "github.com/gin-gonic/gin"

func ProvideBaseHandlers(server *gin.Engine) {
	group := server.Group("api/v1")
	{
		group.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "OK"})
		})
	}
}
