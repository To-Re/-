package Exam

import (
	"bishe/backend/dal"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func ExamList(c *gin.Context) {
	var req ExamListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := ExamListResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateExamListRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	Exams, err := dal.GetExamList()
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询题目列表失败："+err.Error()))
		return
	}

	resExams := make([]*Exam, 0, len(Exams))

	for _, v := range Exams {
		resExams = append(resExams, &Exam{
			ExamId:        int32(v.ID),
			ExamName:      v.Name,
			ExamBeginTime: v.BeginTime.Unix(),
			ExamEndTime:   v.EndTime.Unix(),
			PaperId:       int32(v.PaperID),
		})
	}
	resp.Exams = resExams
	c.JSON(200, resp)
}

type ExamListRequest struct {
}

type ExamListResponse struct {
	Exams []*Exam `json:"exams"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type Exam struct {
	ExamId        int32  `json:"exam_id"`
	ExamName      string `json:"exam_name"`
	ExamBeginTime int64  `json:"exam_begin_time"`
	ExamEndTime   int64  `json:"exam_end_time"`
	PaperId       int32  `json:"paper_id"`
}

func ValidateExamListRequest(req *ExamListRequest) error {
	return nil
}
