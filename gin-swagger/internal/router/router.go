package router

import (
	"gin-swagger/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Engine是gin框架的核心对象，它是路由器，负责处理请求，并调用相应的处理函数。
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/users", handlers.GetUsers)
	}
}

//在这里，我们定义了一个名为SetupRoutes的函数，它接收一个gin.Engine类型的参数，并在这个参数上定义了一组路由。
//在这里，我们定义了一个名为api的组，并在这个组上定义了两个路由：
//GET /api/v1/users，它调用handlers.GetUsers处理函数；
//GET /api/v1/users/:id，它还未定义，但我们可以假设它调用handlers.GetUserByID处理函数。
