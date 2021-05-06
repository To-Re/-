package student

import (
	"bishe/backend/pack"
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

	userId, err := util.GetId(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
		return
	}
	if err = pack.StudentExamCommit(userId, req.ExamId, req.QuestionAnswers); err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "提交失败："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type StudentExamCommitRequest struct {
	ExamId          int32                         `json:"exam_id"`
	QuestionAnswers []*pack.StudentQeustionAnswer `json:"question_answers"`
}

type StudentExamCommitResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateStudentExamCommitRequest(req *StudentExamCommitRequest) error {
	return nil
}
