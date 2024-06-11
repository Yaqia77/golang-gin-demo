package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 权限验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetHeader("User")
		if user == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Set("user", user)
		c.Next()

	}
}
