package question

import (
	"bishe/backend/dal"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func QuestionList(c *gin.Context) {
	var req QuestionListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := QuestionListResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateQuestionListRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	questions, err := dal.GetQuestionList()
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询题目列表失败："+err.Error()))
		return
	}

	resQuestions := make([]*Question, 0, len(questions))

	for _, v := range questions {
		resQuestions = append(resQuestions, &Question{
			QuestionId:     int32(v.ID),
			QuestionDesc:   v.Desc,
			QuestionAnswer: v.Answer,
			QuestionType:   util.QuestionTypeMap[int32(v.Type)],
			OptionDescA:    v.OptionDescA,
			OptionDescB:    v.OptionDescB,
			OptionDescC:    v.OptionDescC,
			OptionDescD:    v.OptionDescD,
		})
	}
	resp.Questions = resQuestions
	c.JSON(200, resp)
}

type QuestionListRequest struct {
}

type QuestionListResponse struct {
	Questions []*Question `json:"questions"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type Question struct {
	QuestionId     int32  `json:"question_id"`
	QuestionDesc   string `json:"question_desc"`
	QuestionAnswer string `json:"question_answer"`
	QuestionType   string `json:"question_type"`
	OptionDescA    string `json:"question_desc_a"`
	OptionDescB    string `json:"question_desc_b"`
	OptionDescC    string `json:"question_desc_c"`
	OptionDescD    string `json:"question_desc_d"`
}

func ValidateQuestionListRequest(req *QuestionListRequest) error {
	return nil
}
