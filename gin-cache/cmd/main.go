package main

import (
	"gin-cache/internal/router"

	"github.com/gin-gonic/gin"
)

//第三方库 gin-contrib/cache 来实现缓存
// go get -u github.com/gin-contrib/cache
//缓存的效果
//第一次请求会有2秒钟的延迟（模拟数据处理时间），而后续请求在缓存有效期内（1分钟）会立即返回相同的数据，因为它们从缓存中获取数据。
//通过这种方式，可以在gin项目中使用缓存来提高性能

func main() {
	r := gin.Default()

	router.SetRouter(r)

	r.Run(":8080")
}
