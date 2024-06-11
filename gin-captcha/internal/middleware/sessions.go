package middleware

//第三方库session
//go get -u github.com/gin-contrib/sessions

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取session
		sessions := sessions.Default(c)
		//获取session中的username
		captchaID := sessions.Get("captchaID")
		if captchaID == nil {
			//如果session中没有username，则返回401
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
	}
}
