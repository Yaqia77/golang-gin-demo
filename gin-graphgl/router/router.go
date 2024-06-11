package router

import (
    "github.com/gin-gonic/gin"
    "gin-graphgl/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/query", controllers.GraphqlHandler())
    r.GET("/playground", controllers.PlaygroundHandler())

    return r
}
