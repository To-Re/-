package method

import (
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetId(c *gin.Context) (int32, error) {
	if uid, ok := c.Get(util.IDKEY); ok && uid.(int32) > 0 {
		return uid.(int32), nil
	}
	return 0, fmt.Errorf("未登录")
}

func GetWho(c *gin.Context) (int32, error) {
	if who, ok := c.Get(util.WHOKEY); ok && who.(int32) > 0 {
		return who.(int32), nil
	}
	return 0, fmt.Errorf("未登录")
}

// 中间件
func AccessTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authtoken := c.Request.Header.Get("AuthToken")
		if authtoken == "" {
			c.JSON(200, util.BuildError(util.NOTLOGIN, util.ErrMap[util.NOTLOGIN]))
			c.Abort()
			return
		} else {
			userInfo, err := util.UserInfo(authtoken)
			if err != nil || userInfo[util.IDKEY] <= 0 || userInfo[util.WHOKEY] <= 0 {
				c.JSON(200, util.BuildError(util.NOTLOGIN, util.ErrMap[util.NOTLOGIN]))
				c.Abort()
				return
			}
			c.Set(util.IDKEY, userInfo[util.IDKEY])
			c.Set(util.WHOKEY, userInfo[util.WHOKEY])
		}
		c.Next()
	}
}

// 中间件
func CheckTeacherAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, err := GetWho(c)
		if err != nil {
			c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
			return
		}
		if userType != util.UserTypeTeacher {
			c.JSON(200, util.BuildError(util.NOTAUTH, util.ErrMap[util.NOTAUTH]+": 需要老师权限"))
		}
		c.Next()
	}
}
