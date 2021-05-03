package paper

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PaperUpdate(c *gin.Context) {
	var req PaperUpdateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := PaperUpdateResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidatePaperUpdateRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	if err := dal.UpdatePaperName(&model.Paper{
		ID:   int(req.PaperId),
		Name: req.PaperName,
	}); err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "修改失败："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type PaperUpdateRequest struct {
	PaperId   int32  `json:"paper_id"`
	PaperName string `json:"paper_name"`
}

type PaperUpdateResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidatePaperUpdateRequest(req *PaperUpdateRequest) error {
	if req.PaperId <= 0 {
		return fmt.Errorf("错误的题目 ID")
	}
	if req.PaperName == "" {
		return fmt.Errorf("考卷名称不得为空")
	}
	return nil
}
