package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": "User" + id,
	})
}

func SearchUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"users": []string{
			"User1",
			"User2",
			"User3",
		},
	})
}

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	emain := c.PostForm("email")
	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"email": emain,
	})

}
