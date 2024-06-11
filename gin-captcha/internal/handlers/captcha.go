package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 初始化 captcha 包，设置验证码存活时间
func init() {
	// 设置验证码存活时间为10分钟
	//SetCustomStore是设置自定义的存储器，这里使用内存存储器
	captcha.SetCustomStore(captcha.NewMemoryStore(1000, time.Minute*10))
}

// GenerateCaptcha是生成验证码的接口
func GenerateCaptcha(c *gin.Context) {
	// 生成验证码
	length := 6
	// captcha.NewLen(length)生成一个长度为length的验证码
	captchaID := captcha.NewLen(length)

	// 将验证码ID存入session
	// sessions := sessions.Default(c)
	// sessions.Set("captcha_id", captchaID)
	// sessions.Save()

	c.JSON(http.StatusOK, gin.H{
		"captcha_id":  captchaID,
		"captcha_url": "api/v1/captcha/" + captchaID + ".png",
	})

	log.Printf("生成验证码id: %s", captchaID)
}

// CaptchaServe 服务验证码图片
func CaptchaServe(c *gin.Context) {
	// 获取验证码ID
	sessions := sessions.Default(c)
	captchaID := sessions.Get("captcha_id")

	if captchaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "captcha_id is required",
		})
		return
	}
	// 输出日志
	log.Printf("图片验证码ID: %s", captchaID)
	// c.Writer.WriteHeader(http.StatusOK)

	// if !captcha.Reload(captchaID) { // 确保 captchaID 有效
	// 	log.Printf("Captcha ID not found or expired: %s", captchaID)
	// 	c.JSON(http.StatusNotFound, gin.H{"message": "Captcha ID not found"})
	// 	return
	// }

	c.Header("Content-Type", "image/png")
	// 输出验证码图片
	//WriteImage 是将验证码图片写入到http.ResponseWriter中，参数分别为：http.ResponseWriter、验证码ID、图片宽度、图片高度

	err := captcha.WriteImage(c.Writer, captchaID.(string), 240, 80)

	if err != nil {
		log.Printf("Error writing captcha image: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating captcha"})
	}
}

// VerifyCaptcha是验证验证码的接口
func VerifyCaptcha(c *gin.Context) {
	// 绑定请求参数
	var RequestBody struct {
		CaptchaID string `json:"captcha_id" binding:"required"`
		Value     string `json:"value" binding:"required"`
	}
	// 验证参数
	if err := c.ShouldBindJSON(&RequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 验证验证码
	if captcha.VerifyString(RequestBody.CaptchaID, RequestBody.Value) {
		c.JSON(http.StatusOK, gin.H{
			"message": "验证码正确",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误"})
	}
}
