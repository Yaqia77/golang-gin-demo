package controllers

import (
	"fmt"
	"gin-blogs/models"
	"gin-blogs/utils"

	"gin-blogs/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	fmt.Println("111111111", user)
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("error: ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("222222", user)
	hashedPassword, err := utils.HashPassword(user.Password)

	fmt.Println("33333333", hashedPassword)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = string(hashedPassword)
	result, err := service.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	// fmt.Println("result1111 :", result)
	utils.SuccessResponse(c, result)
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := service.GetUser(user.Username, user.Password)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	utils.SuccessResponse(c, result)
}
