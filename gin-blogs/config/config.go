package config

import (
	"context"
	"fmt"
	"gin-blogs/models"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
	Ctx = context.Background()
)

func InitDB() {
	// connect to database here
	dsn := "root:123456@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("11111111111112222221", err)
		log.Fatalf("Could not connect to the database: %v", err)
	}

	models.AutoMigrateUserTable(DB)

	models.AutoMigrateBlogTable(DB)

	// Redis 连接
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	// 测试 Redis 连接
	_, err = RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis successfully")
}
