package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 创建表 user表
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
