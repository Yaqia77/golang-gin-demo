package main

import (
	_ "gin-swagger/docs" // 重要：导入docs包
	"gin-swagger/internal/router"

	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title MyApp API
// @version 1.0
// @description This is a sample server for MyApp.
// @host localhost:8080
// @BasePath /api/v1


//WrapHandler是gin-swagger提供的函数，用来注册swagger文档。
//swagFiles.Handler是swaggo提供的函数，用来读取swagger文档的json文件。
//ginSwagger.WrapHandler(swagFiles.Handler)将swagFiles.Handler包装成gin-swagger可以识别的函数。
//最后，将注册swagger文档的路由注册到gin的默认路由器中。

func main() {
	r := gin.Default()

	router.SetupRoutes(r)
	// 注册swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))

	r.Run(":8080")
}
