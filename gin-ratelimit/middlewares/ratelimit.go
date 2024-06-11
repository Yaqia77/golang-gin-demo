package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter 限流器
type RateLimiter struct {
	requests map[string]int // 请求次数
	mu       sync.Mutex     // 互斥锁
	limit    int            // 限制次数
	window   time.Duration  // 时间窗口
}

// NewRateLimit 创建限流器 限制次数为limit 每分钟窗口时间为window
func NewRateLimit(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string]int),
		limit:    limit,
		window:   window,
	}
	go rl.cleanup() // 启动清理任务
	return rl       // 返回限流器
}

// cleanup 清理任务
func (rl *RateLimiter) cleanup() {
	for {
		time.Sleep(rl.window)              // 等待窗口时间
		rl.mu.Lock()                       // 加锁
		rl.requests = make(map[string]int) // 清空请求次数
		rl.mu.Unlock()                     // 解锁

	}

}

// Allow 允许访问
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock() 
	// 记录请求次数
	rl.requests[ip]++
	// 判断是否超过限制
	if rl.requests[ip] > rl.limit {
		return false
	}
	return true

}

func RateLimitMiddleware() gin.HandlerFunc {
	rl := NewRateLimit(5, time.Minute)
	return func(c *gin.Context) {
		// 限制IP每分钟最多5个请求
		ip := c.ClientIP()
		// 超过限制
		if !rl.Allow(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}

}
