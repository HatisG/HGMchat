package pkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 密钥
var jwtKey = []byte("my-secret-key")

// 存储用户信息
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// 生成token，登录成功调用
func GenerateToken(userID uint) (string, error) {
	//过期时间24h
	expirationTime := time.Now().Add(24 * time.Hour)

	//构造payload
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	//签名生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)

}

// 解析token，接口鉴权
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if err != nil || token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
