package handlers

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)



func GetData(c *gin.Context)  {
	//模拟数据处理
	time.Sleep(2 * time.Second)
	//data 模拟的数据
	data := map[string]interface{}{
		"message":"This is some data",
		"timestamp":time.Now().Unix(),
	}

	c.JSON(http.StatusOK, data)
}