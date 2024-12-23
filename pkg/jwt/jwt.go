package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("IAMLEIzZ")

type MyClaims struct {
	UserID int64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenRegisteredClaims 使用默认声明创建jwt
func GenToken(userID int64, username string) (string, error) {
	// 创建 Claims
	claims := MyClaims{
		UserID: userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(MySecret)
}

// ParseRegisteredClaims 解析jwt
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil { // 解析token失败
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid error")
}
