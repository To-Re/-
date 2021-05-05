package student

import (
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func StudentExamCommit(c *gin.Context) {
	var req StudentExamCommitRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := StudentExamCommitResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateStudentExamCommitRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type StudentExamCommitRequest struct {
	ExamId          int32                    `json:"exam_id"`
	QuestionAnswers []*StudentQeustionAnswer `json:"question_answers"`
}

type StudentExamCommitResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateStudentExamCommitRequest(req *StudentExamCommitRequest) error {
	return nil
}

type StudentQeustionAnswer struct {
	QuesitonId int32  `json:"question_id"`
	Answer     string `json:"answer"`
}
