package utils

import (
	"im/common/imerr"
	"time"

	"github.com/golang-jwt/jwt"
)

// 私钥
var secretKey = []byte("peerwise")

// 自定义有效载荷(这里采用自定义的Name和Email作为有效载荷的一部分)
type Claims struct {
	ID	uint	`json:"id"`
	Name	string	`json:"name"`
	// StandardClaims结构体实现了Claims接口(Valid()函数)
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, name string) (string, error) {
	claims := Claims{
		ID: id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24*time.Hour).Unix(),
			Issuer: "im",
		},
	}
	// 指定编码的算法为jwt.SigningMethodHS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// 验证token
func ParseToken(tokenStr string) (*Claims, error) {
	// 输入用户自定义的Claims结构体对象,token,以及自定义函数来解析token字符串为jwt的Token结构体指针
	// Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if token != nil {
		// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	} else {
		// 此时 err != nil
		// jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, imerr.ExpiredTokenErr
			}
		}
	}

	return nil, imerr.InvaildTokenErr
}