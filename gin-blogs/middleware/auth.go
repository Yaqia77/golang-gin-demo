package middleware

import (
	"fmt"
	"gin-blogs/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		fmt.Println("token111111111", authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is missing",
			})
			c.Abort()
			return
		}
		//TrimPrefix是去掉字符串开头的指定内容
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		//验证token是否有效
		claims, err := utils.ValidateJWT(tokenString)
		log.Println("claims:", claims)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("id", claims)
		c.Next()
	}
}
