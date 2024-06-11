package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RBACMiddleware 角色访问控制中间件
func RBACMiddleware(requiredRole string) gin.HandlerFunc {

	return func(c *gin.Context) {
		role, exists := c.Get("role")
		// 检查角色是否存在并且是否等于所需的角色
		if !exists || role != requiredRole {
			// 角色不存在或者角色不匹配
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
