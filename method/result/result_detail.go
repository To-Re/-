package result

import (
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func StudentResultList(c *gin.Context) {
	var req StudentResultListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := StudentResultListResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateStudentResultListRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	info, err := pack.StudentResultList(req.ExamId, req.KlassId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询失败："+err.Error()))
		return
	}

	resp.StudentResult = info
	c.JSON(200, resp)
}

type StudentResultListRequest struct {
	ExamId  int32 `json:"exam_id" form:"exam_id"`
	KlassId int32 `json:"klass_id" form:"klass_id"`
}

type StudentResultListResponse struct {
	StudentResult []*pack.StudentResult `json:"student_results"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateStudentResultListRequest(req *StudentResultListRequest) error {
	return nil
}
