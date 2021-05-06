package result

import (
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func GetResultList(c *gin.Context) {
	var req GetResultListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := GetResultListResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateGetResultListRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	info, err := pack.GetResultList()
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "获取信息失败："+err.Error()))
		return
	}

	resp.Results = info
	c.JSON(200, resp)
}

type GetResultListRequest struct {
}

type GetResultListResponse struct {
	Results []*pack.Result `json:"results"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateGetResultListRequest(req *GetResultListRequest) error {
	return nil
}
