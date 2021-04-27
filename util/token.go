package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "114514" //私钥
	IDKEY     = "id"
	WHOKEY    = "who"
)

func CreteToken(id int32, who int32) (string, error) {
	//生成token
	maxAge := 60 * 60 * 24
	//或者用下面自定义claim
	claims := jwt.MapClaims{
		IDKEY:  id,
		WHOKEY: who,
		"exp":  time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRETKEY))
}

//解析token
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func UserInfo(tokenString string) (map[string]int32, error) {
	ret, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return map[string]int32{
		IDKEY:  ret[IDKEY].(int32),
		WHOKEY: ret[WHOKEY].(int32),
	}, nil
}
