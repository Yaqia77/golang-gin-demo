package controllers

import (
	"gin-blogs/models"
	"gin-blogs/service"
	"gin-blogs/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllBlogs(c *gin.Context) {
	blog, err := service.GetAllBlogs()
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResponse(c, blog)
}

func GetBlogById(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	blog, err := service.GetBlogById(intID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResponse(c, blog)
}

func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	blog.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	blog.UpdatedAt = blog.CreatedAt
	blogs, err := service.CreateBlog(blog)
	if err != nil {
		utils.ErrorResponse(c, http.StatusConflict, err.Error())
		return
	}
	utils.SuccessResponse(c, blogs)

}

func UpdateBlog(c *gin.Context) {
	var blog models.Blog
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := c.ShouldBindJSON(&blog); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	blog.ID = intID
	blog.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	blogs, err := service.UpdateBlog(blog)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResponse(c, blogs)

}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	err = service.DeleteBlog(intID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResponse(c, nil)
}
