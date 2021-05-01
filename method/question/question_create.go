package question

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func QuestionCreate(c *gin.Context) {
	var req QuestionCreateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := QuestionCreateResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateQuestionCreateRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	if err := dal.CreateQuestion(&model.Question{
		Desc:        req.QuestionDesc,
		Type:        int(req.QuestionType),
		Answer:      req.QuestionAnswer,
		OptionDescA: req.OptionDescA,
		OptionDescB: req.OptionDescB,
		OptionDescC: req.OptionDescC,
		OptionDescD: req.OptionDescD,
	}); err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "创建题目："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type QuestionCreateRequest struct {
	QuestionDesc   string `json:"question_desc"`
	QuestionAnswer string `json:"question_answer"`
	QuestionType   int32  `json:"question_type"`
	OptionDescA    string `json:"option_desc_a"`
	OptionDescB    string `json:"option_desc_b"`
	OptionDescC    string `json:"option_desc_c"`
	OptionDescD    string `json:"option_desc_d"`
}

type QuestionCreateResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateQuestionCreateRequest(req *QuestionCreateRequest) error {
	if req.QuestionAnswer == "" {
		return fmt.Errorf("答案不能为空")
	}
	if req.QuestionDesc == "" {
		return fmt.Errorf("题目描述不能为空")
	}
	if req.QuestionType <= 0 {
		return fmt.Errorf("题目类型不能为空")
	}
	if req.OptionDescA == "" || req.OptionDescB == "" || req.OptionDescC == "" || req.OptionDescD == "" {
		return fmt.Errorf("选项描述不能为空")
	}
	return nil
}
