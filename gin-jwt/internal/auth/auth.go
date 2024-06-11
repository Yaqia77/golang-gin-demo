package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 是一个中间件，用于处理JWT验证
type AuthMiddleware struct {
	jwtAuth *JWTAuth
}

// NewAuthMiddleware 创建AuthMiddleware实例
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		jwtAuth: NewJWTAuth(),
	}
}

// LoginHandler 生成JWT
func (am *AuthMiddleware) LoginHandler(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	token, err := am.jwtAuth.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// AuthMiddleware 将两个方法绑定到gin框架的路由上
func (am *AuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := am.jwtAuth.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
