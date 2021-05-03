package exam

import (
	"bishe/backend/dal"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func ExamDetail(c *gin.Context) {
	var req ExamDetailRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := ExamDetailResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateExamDetailRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	examInfo, err := dal.GetExamById(req.ExamId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询失败："+err.Error()))
		return
	}

	resp.ExamName = examInfo.Name
	resp.ExamBeginTime = examInfo.BeginTime.Unix()
	resp.ExamEndTime = examInfo.EndTime.Unix()
	resp.PaperId = int32(examInfo.PaperID)

	c.JSON(200, resp)
}

type ExamDetailRequest struct {
	ExamId int32 `form:"exam_id"`
}

type ExamDetailResponse struct {
	ExamName      string `json:"exam_name"`
	ExamBeginTime int64  `json:"exam_begin_time"`
	ExamEndTime   int64  `json:"exam_end_time"`
	PaperId       int32  `json:"paper_id"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateExamDetailRequest(req *ExamDetailRequest) error {
	return nil
}
