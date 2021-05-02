package paper

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PaperCreate(c *gin.Context) {
	var req PaperCreateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := PaperCreateResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidatePaperCreateRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	if err := dal.CreatePaper(&model.Paper{
		Name: req.PaperName,
	}); err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "创建题目："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type PaperCreateRequest struct {
	PaperName string `json:"paper_name"`
}

type PaperCreateResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidatePaperCreateRequest(req *PaperCreateRequest) error {
	if req.PaperName == "" {
		return fmt.Errorf("试卷名不能为空")
	}
	return nil
}
