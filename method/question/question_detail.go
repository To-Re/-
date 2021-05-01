package question

import (
	"bishe/backend/dal"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func QuestionDetail(c *gin.Context) {
	var req QuestionDetailRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := QuestionDetailResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateQuestionDetailRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	QuestionInfo, err := dal.GetQuestionById(req.QuestionId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询失败："+err.Error()))
		return
	}

	resp.QuestionDesc = QuestionInfo.Desc
	resp.QuestionAnswer = QuestionInfo.Answer
	resp.QuestionType = int32(QuestionInfo.Type)
	resp.OptionDescA = QuestionInfo.OptionDescA
	resp.OptionDescB = QuestionInfo.OptionDescB
	resp.OptionDescC = QuestionInfo.OptionDescC
	resp.OptionDescD = QuestionInfo.OptionDescD
	c.JSON(200, resp)
}

type QuestionDetailRequest struct {
	QuestionId int32 `json:"Question_id" form:"question_id"`
}

type QuestionDetailResponse struct {
	QuestionDesc   string `json:"question_desc"`
	QuestionAnswer string `json:"question_answer"`
	QuestionType   int32  `json:"question_type"`
	OptionDescA    string `json:"option_desc_a"`
	OptionDescB    string `json:"option_desc_b"`
	OptionDescC    string `json:"option_desc_c"`
	OptionDescD    string `json:"option_desc_d"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateQuestionDetailRequest(req *QuestionDetailRequest) error {
	return nil
}
