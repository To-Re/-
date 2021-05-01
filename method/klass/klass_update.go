package klass

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func KlassUpdate(c *gin.Context) {
	var req KlassUpdateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := KlassUpdateResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateKlassUpdateRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	if err := dal.UpdateKlass(&model.Klass{
		Name: req.KlassName,
		ID:   int(req.KlassId),
	}); err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "修改失败："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type KlassUpdateRequest struct {
	KlassId   int32  `json:"klass_id"`
	KlassName string `json:"klass_name"`
}

type KlassUpdateResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateKlassUpdateRequest(req *KlassUpdateRequest) error {
	if req.KlassId <= 0 {
		return fmt.Errorf("错误的班级 ID")
	}
	if req.KlassName == "" {
		return fmt.Errorf("班级名不得为空")
	}
	return nil
}
