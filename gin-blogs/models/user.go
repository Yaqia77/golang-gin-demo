package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primarykey;autoIncrement"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id" `
	Username string `json:"username" `
	Token    string `json:"token"`
}

func AutoMigrateUserTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}, &UserResponse{}); err != nil {
		return err
	}
	return nil
}
