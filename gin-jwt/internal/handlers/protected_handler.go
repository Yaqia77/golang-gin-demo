package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProtectedHandler 处理保护的请求
func ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, you are authorized"})
}
