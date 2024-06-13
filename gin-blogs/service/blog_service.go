package service

import (
	"gin-blogs/config"
	"gin-blogs/models"
)

func GetAllBlogs() ([]models.Blog, error) {
	var blog []models.Blog

	err := config.DB.Find(&blog).Error

	if err != nil {
		// 返回其他错误
		return []models.Blog{}, err
	}

	return blog, nil
}

func GetBlogById(id int) (models.Blog, error) {
	var blog models.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		return models.Blog{}, err
	}

	return blog, nil
}

func CreateBlog(blog models.Blog) (models.Blog, error) {
	if err := config.DB.Create(&blog).Error; err != nil {
		return models.Blog{}, err
	}

	return blog, nil
}

func UpdateBlog(blog models.Blog) (models.Blog, error) {
	if err := config.DB.Save(&blog).Error; err != nil {
		return models.Blog{}, err
	}

	return blog, nil
}

func DeleteBlog(id int) error {
	if err := config.DB.Delete(&models.Blog{}, id).Error; err != nil {
		return err
	}

	return nil
}
