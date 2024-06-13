package service

import (
	"errors"
	"gin-blogs/config"
	"gin-blogs/models"
	"gin-blogs/utils"

	"gorm.io/gorm"
)

func CreateUser(user models.User) (models.User, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUser(username, password string) (models.UserResponse, error) {
	var user models.User
	var userResult models.UserResponse
	if err := config.DB.Where("username =? AND password = ?", username, utils.CheckPasswordHash(password, user.Password)).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.UserResponse{}, errors.New("找不到该用户")
		}
		return models.UserResponse{}, err
	}

	token, err := utils.GenerateJWT(user.Username)

	userResult.ID = user.ID
	userResult.Username = user.Username
	userResult.Token = token

	if err != nil {
		return models.UserResponse{}, err
	}

	return userResult, nil
}
