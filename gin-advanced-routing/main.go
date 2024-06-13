package main

import (
	"gin-advanced-routing/config"
	"gin-advanced-routing/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//连接数据库的函数
	config.ConnectDatabase()

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", controllers.CreateUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	r.Run() // listen and serv	e on 0.0.0.0:8080
}
