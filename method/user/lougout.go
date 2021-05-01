package user

import (
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	var req LogoutRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := LogoutResponse{
		Code: 0,
		Msg:  "",
	}
	c.JSON(200, resp)
}

type LogoutRequest struct {
}
type LogoutResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
