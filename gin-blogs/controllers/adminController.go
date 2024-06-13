package controllers

// import (
// 	"gin-blogs/models"
// 	"gin-blogs/service"
// 	"gin-blogs/utils"
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// 	"strconv"
// 	"time"
// )

// func GetPost(c *gin.Context) {
// 	id := c.Param("id")
// 	num, err := strconv.Atoi(id)
// 	if err != nil {
// 		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
// 		return
// 	}
// 	post, err := service.GetPost(num)
// 	if err != nil {
// 		utils.ErrorResponse(c, http.StatusNotFound, "Post not found")
// 		return
// 	}
// 	utils.SuccessResponse(c, post)
// }

// func CreatePost(c *gin.Context) {
// 	var post models.Post

// 	if err := c.ShouldBindJSON(&post); err != nil {
// 		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
// 		return
// 	}

// 	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
// 	post.UpdatedAt = post.CreatedAt

// 	result, err := service.CreatePost(post)
// 	if err != nil {
// 		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create post")
// 		return
// 	}

// 	utils.SuccessResponse(c, result)
// }

// func GetAllPosts(c *gin.Context) {
// 	post, err := service.GetAllPosts()
// 	if err != nil {
// 		utils.ErrorResponse(c, http.StatusNotFound, "Post not found")
// 	}
// 	utils.SuccessResponse(c, post)
// }

// func UpdatePost(c *gin.Context) {
// 	id := c.Param("id")

// 	num, err := strconv.Atoi(id)
// 	if err != nil {
// 		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
// 		return
// 	}

// 	var post models.Post
// 	if err := c.ShouldBindJSON(&post); err != nil {
// 		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
// 		return
// 	}

// 	post.ID = num
// 	post.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

// 	result, err := service.UpdatePost(post)
// 	if err != nil {
// 		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update post")
// 		return
// 	}
// 	utils.SuccessResponse(c, result)
// }

// func DeletePost(c *gin.Context) {
// 	id := c.Param("id")
// 	num, err := strconv.Atoi(id)
// 	if err != nil {
// 		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
// 		return
// 	}
// 	err = service.DeletePost(num)
// 	if err != nil {
// 		utils.ErrorResponse(c, http.StatusNotFound, "Post not found")
// 		return
// 	}
// 	utils.SuccessResponse(c, nil)
// }