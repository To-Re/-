package klass

import (
	"bishe/backend/pack"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func KlassCreate(c *gin.Context) {
	var req KlassCreateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := KlassCreateResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateKlassCreateRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}
	userId, err := util.GetId(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, util.ErrMap[util.PARAMERROR]))
	}

	if err := pack.CreateKlassAndBindTeacher(userId, req.KlassName); err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "创建失败："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type KlassCreateRequest struct {
	KlassName string `json:"klass_name"`
}

type KlassCreateResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateKlassCreateRequest(req *KlassCreateRequest) error {
	if req.KlassName == "" {
		return fmt.Errorf("班级名不得为空")
	}
	return nil
}
