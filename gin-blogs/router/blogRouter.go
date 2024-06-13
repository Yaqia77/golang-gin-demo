package router

import (
	"gin-blogs/controllers"
	"gin-blogs/middleware"

	"github.com/gin-gonic/gin"
)

func BlogRouter(c *gin.RouterGroup) {
	AuthMiddleware := middleware.AuthMiddleware()
	blogGroup := c.Group("/blog")
	{
		blogGroup.GET("/", controllers.GetAllBlogs)
		blogGroup.GET("/:id", controllers.GetBlogById)
		blogGroup.POST("/", AuthMiddleware, controllers.CreateBlog)
		blogGroup.PUT("/:id", AuthMiddleware, controllers.UpdateBlog)
		blogGroup.DELETE("/:id", AuthMiddleware, controllers.DeleteBlog)
	}
}
