package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	jwtSigningKey = []byte("secret")
)

// JWTAuth 实现JWT认证

type JWTAuth struct{}

// NewJWTAuth 创建JWTAuth实例
func NewJWTAuth() *JWTAuth {
	return &JWTAuth{}
}

// GenerateToken 生成JWT token
func (ja *JWTAuth) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken 验证JWT token
func (ja *JWTAuth) VerifyToken(c *gin.Context) error {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return errors.New("authorization token is required")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSigningKey, nil
	})
	if err != nil || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
