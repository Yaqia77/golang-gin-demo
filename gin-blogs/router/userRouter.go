package router

import (
	"gin-blogs/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(c *gin.RouterGroup) {
	userGroup := c.Group("/")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}
}
