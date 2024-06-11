package router

import (
	"gin-cache/internal/handlers"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	// 设置缓存存储
	//persistence 是缓存持久化的接口，这里使用的是内存存储
	store := persistence.NewInMemoryStore(time.Second)

	api := r.Group("/api/v1")

	{
		// 使用缓存中间件
		api.GET("/data", cache.CachePage(store, time.Minute, handlers.GetData))
	}
}
