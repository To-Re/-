package student

import (
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func StudentExamDetail(c *gin.Context) {
	var req StudentExamDetailRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := StudentExamDetailResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateStudentExamDetailRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	userId, err := util.GetId(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
		return
	}
	info, err := pack.GetStudentExamDetail(userId, req.ExamId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询失败："+err.Error()))
		return
	}

	resp.ExamName = info.ExamName
	resp.ScoreLimit = info.ScoreLimit
	resp.ExamEndTime = info.ExamEndTime
	resp.Questions = info.Questions
	c.JSON(200, resp)
}

type StudentExamDetailRequest struct {
	ExamId int32 `json:"exam_id" form:"exam_id"`
}

type StudentExamDetailResponse struct {
	ExamName    string           `json:"question_desc"`
	ScoreLimit  int32            `json:"score_limit"`
	ExamEndTime int64            `json:"exam_end_time"`
	Questions   []*pack.Question `json:"questions"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateStudentExamDetailRequest(req *StudentExamDetailRequest) error {
	return nil
}
