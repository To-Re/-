package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID int32 `json:"user_id"`
	Who    int32 `json:"who"`
}

const (
	IDKEY      = "user_id"
	WHOKEY     = "who"
	SECRETKEY  = "114514" //私钥
	ExpireTime = time.Hour * 24
)

func CreteToken(id int32, who int32) (string, error) {
	claims := &JWTClaims{
		UserID: id,
		Who:    who,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Duration(ExpireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRETKEY))
}

//解析token
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, fmt.Errorf("jwt 解析错误")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, fmt.Errorf("jwt 解析错误")
	}
	return claims, nil
}

func UserInfo(tokenString string) (map[string]int32, error) {
	ret, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return map[string]int32{
		IDKEY:  ret.UserID,
		WHOKEY: ret.Who,
	}, nil
}
