package student

import (
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func GetStudentExamList(c *gin.Context) {
	var req StudentExamListUpdateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := StudentExamListUpdateResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateStudentExamListUpdateRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	userId, err := util.GetId(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
		return
	}

	info, err := pack.GetStudentExamList(userId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询失败："+err.Error()))
		return
	}

	res := make([]*Exam, 0, len(info))
	for _, v := range info {
		res = append(res, &Exam{
			ExamId:           v.ExamId,
			ExamName:         v.ExamName,
			ExamBeginTime:    v.ExamBeginTime,
			ExamEndTime:      v.ExamEndTime,
			ExamStudentScore: v.ExamStudentScore,
		})
	}
	resp.Exams = res
	c.JSON(200, resp)
}

type StudentExamListUpdateRequest struct {
}

type StudentExamListUpdateResponse struct {
	Exams []*Exam `json:"exams"`
	Code  int32   `json:"code"`
	Msg   string  `json:"msg"`
}

type Exam struct {
	ExamId           int32  `json:"exam_id"`
	ExamName         string `json:"exam_name"`
	ExamBeginTime    int64  `json:"exam_begin_time"`
	ExamEndTime      int64  `json:"exam_end_time"`
	ExamStudentScore *int32 `json:"exam_student_score,omitempty"`
}

func ValidateStudentExamListUpdateRequest(req *StudentExamListUpdateRequest) error {
	return nil
}
