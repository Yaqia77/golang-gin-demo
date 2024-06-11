package router

import (
	"project02/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", service.RegisterUser)
		userGroup.POST("/create", service.CreateUser)
	}
	return r
}
