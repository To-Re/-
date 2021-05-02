package paper

import (
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func PaperQuestionList(c *gin.Context) {
	var req PaperQuestionListlRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := PaperQuestionListlResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidatePaperQuestionListlRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	paperQuestions, err := pack.GetPaperQuestionList(req.PaperId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询考卷试题列表失败："+err.Error()))
		return
	}

	resPaperQuestions := make([]*PaperQuestion, 0, len(paperQuestions))
	for _, v := range paperQuestions {
		resPaperQuestions = append(resPaperQuestions, &PaperQuestion{
			QuestionId:    v.QuestionId,
			QuestionScore: v.QuestionScore,
			QuestionDesc:  v.QuestionDesc,
			QuestionType:  v.QuestionType,
		})
	}
	resp.PaperQuestions = resPaperQuestions

	c.JSON(200, resp)
}

type PaperQuestionListlRequest struct {
	PaperId int32 `form:"paper_id"`
}

type PaperQuestionListlResponse struct {
	PaperQuestions []*PaperQuestion `json:"paper_questions"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type PaperQuestion struct {
	QuestionId    int32  `json:"question_id"`
	QuestionScore int32  `json:"question_score"`
	QuestionDesc  string `json:"question_desc"`
	QuestionType  string `json:"question_type"`
}

func ValidatePaperQuestionListlRequest(req *PaperQuestionListlRequest) error {
	return nil
}
