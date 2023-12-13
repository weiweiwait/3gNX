package middles

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("your-secret-key")

// 生成Jwt令牌
func GenerateJWT(username string) string {
	// 创建一个新的令牌声明
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为24小时
	}

	// 创建令牌对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对令牌进行签名
	tokenString, _ := token.SignedString(jwtSecret)

	return tokenString
}
