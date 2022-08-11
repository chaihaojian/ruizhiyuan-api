package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//TokenExpireDuration token过期时间
const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("ruizhiyuan")

//MyClaims token所包含的信息
type MyClaims struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userid int64, username string) (string, error) {
	// 创建一个自己的声明
	c := MyClaims{
		userid,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "ruizhiyuan",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
