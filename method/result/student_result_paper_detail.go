package result

import (
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func GetStudentResultPaperDetail(c *gin.Context) {
	var req GetStudentResultPaperDetailRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := GetStudentResultPaperDetailResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateGetStudentResultPaperDetailRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	info, err := pack.GetStudentResultPaperDetail(req.StudentId, req.ExamId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询失败："+err.Error()))
		return
	}

	resp.StudentScore = info.StudentScore
	resp.ExamName = info.ExamName
	resp.ScoreLimit = info.ScoreLimit
	resp.ExamEndTime = info.ExamEndTime
	resp.Questions = info.Questions
	c.JSON(200, resp)
}

type GetStudentResultPaperDetailRequest struct {
	StudentId int32 `json:"student_id" form:"student_id"`
	ExamId    int32 `json:"exam_id" form:"exam_id"`
}

type GetStudentResultPaperDetailResponse struct {
	ExamName     string           `json:"exam_name"`
	ScoreLimit   int32            `json:"score_limit"`
	ExamEndTime  int64            `json:"exam_end_time"`
	StudentScore *int32           `json:"student_score,omitempty"`
	Questions    []*pack.Question `json:"questions"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateGetStudentResultPaperDetailRequest(req *GetStudentResultPaperDetailRequest) error {
	return nil
}
