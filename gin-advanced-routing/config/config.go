package config

import (
	"gin-advanced-routing/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )

	dsn := "root:123456@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 为什么加了：后就会报错：invalid memory address or nil pointer dereference
	// 因为DB没有初始化，所以会报错

	// 解决方法：在ConnectDatabase()函数中初始化DB变量，并调用AutoMigrate()方法

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// 迁移模型
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

}
