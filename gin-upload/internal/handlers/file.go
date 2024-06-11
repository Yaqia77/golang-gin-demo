package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	//获取上传的文件
	file, err := c.FormFile("file")
	//判断文件是否为空
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//上传文件
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		//上传失败
		//StatusInternalServerError 服务器内部错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": file.Filename,
	})
}
