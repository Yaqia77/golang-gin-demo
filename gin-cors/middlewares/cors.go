package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//AllowOrigins是允许跨域请求的源，可以是字符串或者字符串数组，默认是["*"]，表示允许所有源
//AllowMethods是允许跨域请求的方法，可以是字符串或者字符串数组，默认是["GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"]
//AllowHeaders是允许跨域请求的头部，可以是字符串或者字符串数组，默认是["Origin", "Content-Type"]
//ExposeHeaders是允许跨域请求时返回的头部，可以是字符串或者字符串数组，默认是["Content-Length"]
//AllowCredentials是允许跨域请求时携带cookie，默认是false
//MaxAge是允许跨域请求的最大时长，默认是0，表示不限制

func CORSMiddleware() gin.HandlerFunc {

	//cors.Config是cors中间件的配置
	return cors.New(cors.Config{
		
		AllowOrigins:     []string{"https://example.com", "https://anotherexample.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours
	})
}