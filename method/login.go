package method

import (
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := LoginResponse{
		Code: 0,
		Msg:  "",
	}
	// 校验
	// userId 从数据库拿到，通过number
	userId := 123
	ret, _ := util.CreteToken(int32(userId), int32(req.UserType))
	fmt.Println("token : " + ret)
	resp.Token = ret
	c.JSON(200, resp)
}

type LoginRequest struct {
	Number   string `json:"number"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
}
type LoginResponse struct {
	Token string `json:"token"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
