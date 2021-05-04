package exam

import (
	"bishe/backend/model"
	"bishe/backend/pack"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ExamKlassBind(c *gin.Context) {
	var req ExamKlassBindRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := ExamKlassBindResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateExamKlassBindRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	err = pack.ExamKlassBind(&model.KlassExam{
		KlassID: int(req.KlassID),
		ExamID:  int(req.ExamId),
	})
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "创建失败："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type ExamKlassBindRequest struct {
	ExamId  int32 `json:"exam_id"`
	KlassID int32 `json:"klass_id"`
}

type ExamKlassBindResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateExamKlassBindRequest(req *ExamKlassBindRequest) error {
	if req.ExamId <= 0 {
		return fmt.Errorf("错误的考试id")
	}
	if req.KlassID <= 0 {
		return fmt.Errorf("错误的班级id")
	}
	return nil
}
