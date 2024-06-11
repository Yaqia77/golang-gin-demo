package main

import (
	"gin-captcha/internal/router"
	// "os"

	// "github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

//使用验证码可以通过第三方库github.com/dchest/captcha来实现
//gin-captcha项目中使用了github.com/gin-gonic/gin框架，gin框架是一个轻量级的Go Web框架，使用起来非常方便。
//下载
//go get -u github.com/gin-gonic/gin
//go get -u github.com/dchest/captcha

func main() {
	// captchaID := captcha.NewLen(6)
	// file, err := os.Create("captcha.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// err = captcha.WriteImage(file, captchaID, captcha.StdWidth, captcha.StdHeight)
	// if err != nil {
	// 	panic(err)
	// }

	// println("Captcha saved to captcha.png")

	r := gin.Default()

	router.SetupRoutes(r)

	r.Run(":8080")
}
