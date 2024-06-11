package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//获取session
	session := sessions.Default(c)
	//从session中获取username
	username := c.PostForm("username")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}
	//设置session中的username
	session.Set("username", username)
	//保存session
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Login in successful"})
}

func Logout(c *gin.Context) {
	//获取session
	session := sessions.Default(c)
	//清除session
	session.Clear()
	//保存session
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Logout out successful"})

}

func Profile(c *gin.Context) {
	sessions := sessions.Default(c)

	username := sessions.Get("username")

	if username == nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": username})

}
