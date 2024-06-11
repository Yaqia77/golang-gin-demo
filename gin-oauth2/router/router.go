package router

import (
    "github.com/gin-gonic/gin"
    "gin-oauth2/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/login", controllers.HandleLogin)
    r.GET("/callback", controllers.HandleCallback)

    return r
}
