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

// @Summary 获取所有博客
// @Description 获取所有博客
// @Produce json
// @Success 200 {array} models.Blog
// @Router /api/blog [get]
func GetAllBlogs(c *gin.Context) {
	blog, err := service.GetAllBlogs()
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResponse(c, blog)
}

// GetBlogById godoc
// @Summary Get a blog by ID
// @Description Get a blog by ID
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} models.Blog
// @Router /api/blog/{id} [get]
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

// CreateBlog godoc
// @Summary Create a blog
// @Description Create a new blog
// @Produce json
// @Param blog body models.Blog true "Blog object"
// @Success 201 {object} models.Blog
// @Router /api/blog [post]
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

// UpdateBlog godoc
// @Summary Update a blog
// @Description Update an existing blog
// @Produce json
// @Param id path int true "Blog ID"
// @Param blog body models.Blog true "Blog object"
// @Success 200 {object} models.Blog
// @Router /api/blog/{id} [put]
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

// DeleteBlog godoc
// @Summary Delete a blog
// @Description Delete a blog by ID
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200
// @Router /api/blog/{id} [delete]
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
