package router

import (
	"gin-params/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SerRouter(c *gin.Engine) {
	api := c.Group("/api")
	{
		api.GET("/users/:id", handlers.GetUserByID)
		api.GET("/search", handlers.SearchUser)
		api.POST("/users", handlers.CreateUser)
	}
}
