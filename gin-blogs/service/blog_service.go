package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-blogs/config"
	"gin-blogs/models"
	"time"
)

func GetAllBlogs() ([]models.Blog, error) {
	var blog []models.Blog
	// 从redis中获取数据
	ctx := context.Background()
	key := "all_blogs"
	result, err := config.RDB.Get(ctx, key).Result()
	if err == nil {
		err := json.Unmarshal([]byte(result), &blog)
		if err != nil {
			return []models.Blog{}, nil
		}
		return blog, nil
	}
	// 从数据库中获取数据
	if err := config.DB.Find(&blog).Error; err != nil {
		// 返回其他错误
		return []models.Blog{}, err
	}

	return blog, nil
}

func GetBlogById(id int) (models.Blog, error) {
	var blog models.Blog
	ctx := context.Background()
	key := fmt.Sprintf("blog:%d", id)
	result, err := config.RDB.Get(ctx, key).Result()
	if err == nil {
		err := json.Unmarshal([]byte(result), &blog)
		if err != nil {
			return models.Blog{}, nil
		}
		return blog, nil
	}

	if err := config.DB.First(&blog, id).Error; err != nil {
		return models.Blog{}, err
	}
	blogJSON, err := json.Marshal(blog)
	if err != nil {
		return models.Blog{}, err
	}
	config.RDB.Set(ctx, key, string(blogJSON), time.Hour)
	return blog, nil
}

func CreateBlog(blog models.Blog) (models.Blog, error) {

	if err := config.DB.Create(&blog).Error; err != nil {
		return models.Blog{}, err
	}
	// 更新redis中的数据
	blogJSON, err := json.Marshal(blog)
	if err != nil {
		return models.Blog{}, err
	}
	ctx := context.Background()
	key := fmt.Sprintf("blog:%d", blog.ID)
	err = config.RDB.Set(ctx, key, string(blogJSON), time.Hour).Err()
	if err != nil {
		fmt.Println("Error updating post in Redis:", err)
	}
	return blog, nil
}

func UpdateBlog(blog models.Blog) (models.Blog, error) {
	ctx := context.Background()

	if err := config.DB.Save(&blog).Error; err != nil {
		return models.Blog{}, err
	}
	// 更新redis中的数据
	key := fmt.Sprintf("blog:%d", blog.ID)
	blogJSON, err := json.Marshal(blog)
	if err != nil {
		return models.Blog{}, err
	}
	//Set的参数分别为key，value，过期时间
	err = config.RDB.Set(ctx, key, string(blogJSON), time.Hour).Err()
	if err != nil {
		fmt.Println("Error updating post in Redis:", err)
	}

	return blog, nil
}

func DeleteBlog(id int) error {
	if err := config.DB.Delete(&models.Blog{}, id).Error; err != nil {
		return err
	}
	key := fmt.Sprintf("blog:%d", id)
	config.RDB.Del(context.Background(), key)
	return nil
}
