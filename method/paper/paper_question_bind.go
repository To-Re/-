package paper

import (
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PaperQuestionBind(c *gin.Context) {
	var req PaperQuestionBindRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := PaperQuestionBindResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidatePaperQuestionBindRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type PaperQuestionBindRequest struct {
	PaperId       int32 `json:"paper_id"`
	QuestionId    int32 `json:"question_id"`
	QuestionScore int32 `json:"question_score"`
}

type PaperQuestionBindResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidatePaperQuestionBindRequest(req *PaperQuestionBindRequest) error {
	if req.PaperId <= 0 {
		return fmt.Errorf("错误的考卷id")
	}
	if req.QuestionId <= 0 {
		return fmt.Errorf("错误的考题id")
	}
	if req.QuestionScore < 0 {
		return fmt.Errorf("题目分数不得小于0分")
	}
	return nil
}
