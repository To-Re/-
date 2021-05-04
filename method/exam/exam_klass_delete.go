package exam

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ExamKlassDelete(c *gin.Context) {
	var req ExamKlassDeleteRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := ExamKlassDeleteResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateExamKlassDeleteRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	// 删除
	err = dal.DeleteKlassExam(&model.KlassExam{
		KlassID: int(req.KlassID),
		ExamID:  int(req.ExamId),
	})
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "删除失败："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type ExamKlassDeleteRequest struct {
	ExamId  int32 `json:"exam_id"`
	KlassID int32 `json:"klass_id"`
}

type ExamKlassDeleteResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateExamKlassDeleteRequest(req *ExamKlassDeleteRequest) error {
	if req.ExamId <= 0 {
		return fmt.Errorf("错误的考试id")
	}
	if req.KlassID <= 0 {
		return fmt.Errorf("错误的班级id")
	}
	return nil
}
