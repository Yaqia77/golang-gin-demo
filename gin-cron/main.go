package main

//gin-cron 定时任务集成

import (
	"net/http"
	"time"

	"gin-cron/jobs"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	router := gin.Default()

	// Set up cron jobs
	//创建一个新的cron实例
	c := cron.New()
	//添加一个定时任务，每隔1分钟执行一次jobs.MyCronJob()
	_, err := c.AddFunc("@every 1m", jobs.MyCronJob)
	if err != nil {
		//panic是用于中断程序的函数，这里是因为添加定时任务失败了，程序需要中断
		panic(err)
	}

	_, err = c.AddFunc("@every 2m", jobs.AnotherCronJob)
	if err != nil {
		panic(err)
	}

	c.Start()
	defer c.Stop()

	// Gin routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Gin with Cron!")
	})

	router.GET("/time", func(c *gin.Context) {
		c.String(http.StatusOK, "Current time: %s", time.Now().Format(time.RFC1123))
	})

	router.Run(":8080")
}
