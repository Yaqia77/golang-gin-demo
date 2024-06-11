package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /users [get]
func GetUsers(c *gin.Context) {
	//
	users := []string{"user1", "user2"}
	c.JSON(http.StatusOK, users)
}
