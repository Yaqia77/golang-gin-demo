package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 模拟一个用户数据存储
var users = map[string]string{
	"user1":  "user",
	"admin1": "admin",
}

// AuthMiddleware 验证用户身份

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取用户信息
		user := c.GetHeader("User")
		// 从用户数据存储中获取用户角色
		role, exist := users[user]

		if !exist {
			// 用户不存在，返回
			//StatusUnauthorized 401状态码，表示未授权
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		// 将用户角色存入上下文
		c.Set("role", role)
		// 继续处理请求
		c.Next()
	}
}
