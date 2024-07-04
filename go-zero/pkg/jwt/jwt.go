package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     int    `json:"role"` // 1: user, 2: admin
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

func GenToken(user JwtPayLoad, accessSecret string, expires int64) (string, error) {
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expires)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))

}

func ParseToken(tokenStr string, accessSecret string, expires int64) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if time.Now().Unix() > claims.ExpiresAt {
			return nil, err
		}
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
