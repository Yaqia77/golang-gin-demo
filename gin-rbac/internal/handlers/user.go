package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PublicEndpoint godoc
// @Summary Public endpoint
// @Description Accessible by anyone
// @Tags public
// @Produce json
// @Success 200 {string} string "ok"
// @Router /public [get]
func PublicEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "This is a public endpoint")
}

// UserEndpoint godoc
// @Summary User endpoint
// @Description Accessible by authenticated users
// @Tags user
// @Produce json
// @Success 200 {string} string "ok"
// @Router /user [get]
func UserEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "This is an endpoint for authenticated users")
}

// AdminEndpoint godoc
// @Summary Admin endpoint
// @Description Accessible by admin users
// @Tags admin
// @Produce json
// @Success 200 {string} string "ok"
// @Router /admin [get]

func AdminEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "This is an endpoint for admin users")
}
