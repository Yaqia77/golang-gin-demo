package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	UserRouter(api)

	BlogRouter(api)

	return r

	
}
