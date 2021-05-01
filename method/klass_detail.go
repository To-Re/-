package method

import (
	"bishe/backend/dal"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func KlassDetail(c *gin.Context) {
	var req KlassDetailRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := KlassDetailResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateKlassDetailRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	klassInfo, err := dal.GetKlassById(req.KlassId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询失败："+err.Error()))
		return
	}

	resp.KlassName = klassInfo.Name
	c.JSON(200, resp)
}

type KlassDetailRequest struct {
	KlassId int32 `json:"klass_id" form:"klass_id"`
}

type KlassDetailResponse struct {
	KlassName string `json:"klass_name"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateKlassDetailRequest(req *KlassDetailRequest) error {
	return nil
}
