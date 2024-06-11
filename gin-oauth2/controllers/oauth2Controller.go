package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// 创建OAuth2 配置
// ClientID 和 ClientSecret 是从Google API Console获取的
var (
	// googleOauthConfig 用于Google登录的配置
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauth2StateString = "random"
)

// HandLogin 处理登录请求
func HandleLogin(c *gin.Context) {
	// 重定向到Google登录页面
	//googleOauthConfig.AuthCodeURL是生成授权码的URL,oauth2StateString是用于防止CSRF攻击的随机字符串

	url := googleOauthConfig.AuthCodeURL(oauth2StateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandCallback 处理回调请求
func HandleCallback(c *gin.Context) {
	// state参数用于防止CSRF攻击
	state := c.DefaultQuery("state", "")

	if state != oauth2StateString {
		// 状态不匹配，可能是CSRF攻击
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid oauth state",
		})
		return
	}
	// 获取授权码
	code := c.DefaultQuery("code", "")
	log.Printf("code: %s", code)
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	log.Printf("token: %v", token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "code exchange failed",
			"details": err.Error(),
		})
		return
	}

	// 使用token获取用户信息
	//response是用户信息的响应
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed getting user info",
		})
		return
	}
	defer response.Body.Close()
	// 解析用户信息
	// 这里假设用户信息是JSON格式

	//map是一种无序的键值对的集合，其中每个键都是唯一的，值可以是任意类型。
	var userInfo map[string]interface{}

	//NewDecoder()函数创建一个新的json.Decoder对象，并使用response.
	//Decode()方法将JSON数据解析为map。
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed decoding user info",
		})
		return
	}
	c.JSON(http.StatusOK, userInfo)
}
