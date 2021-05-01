package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetId(c *gin.Context) (int32, error) {
	if uid, ok := c.Get(IDKEY); ok && uid.(int32) > 0 {
		return uid.(int32), nil
	}
	return 0, fmt.Errorf("未登录")
}

func GetWho(c *gin.Context) (int32, error) {
	if who, ok := c.Get(WHOKEY); ok && who.(int32) > 0 {
		return who.(int32), nil
	}
	return 0, fmt.Errorf("未登录")
}

// 中间件
func AccessTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authtoken := c.Request.Header.Get("AuthToken")
		if authtoken == "" {
			c.JSON(200, BuildError(NOTLOGIN, ErrMap[NOTLOGIN]))
			c.Abort()
			return
		} else {
			userInfo, err := UserInfo(authtoken)
			if err != nil || userInfo[IDKEY] <= 0 || userInfo[WHOKEY] <= 0 {
				c.JSON(200, BuildError(NOTLOGIN, ErrMap[NOTLOGIN]))
				c.Abort()
				return
			}
			c.Set(IDKEY, userInfo[IDKEY])
			c.Set(WHOKEY, userInfo[WHOKEY])
		}
		c.Next()
	}
}

// 中间件
func CheckTeacherAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, err := GetWho(c)
		if err != nil {
			c.JSON(200, BuildError(NOTLOGIN, err.Error()))
			c.Abort()
			return
		}
		if userType != UserTypeTeacher {
			c.JSON(200, BuildError(NOTAUTH, ErrMap[NOTAUTH]+": 需要老师权限"))
			c.Abort()
			return
		}
		c.Next()
	}
}
