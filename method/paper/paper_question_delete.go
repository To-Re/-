package paper

import (
	"bishe/backend/model"
	"bishe/backend/pack"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PaperQuestionDelete(c *gin.Context) {
	var req PaperQuestionDeleteRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := PaperQuestionDeleteResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidatePaperQuestionDeleteRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	// 删除
	err = pack.PaperQuestionDelete(&model.PaperQuestion{
		PaperID:    int(req.PaperId),
		QuestionID: int(req.QuestionId),
	})
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "删除失败："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type PaperQuestionDeleteRequest struct {
	PaperId    int32 `json:"paper_id"`
	QuestionId int32 `json:"question_id"`
}

type PaperQuestionDeleteResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidatePaperQuestionDeleteRequest(req *PaperQuestionDeleteRequest) error {
	if req.PaperId <= 0 {
		return fmt.Errorf("错误的考卷id")
	}
	if req.QuestionId <= 0 {
		return fmt.Errorf("错误的考题id")
	}
	return nil
}
