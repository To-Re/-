package exam

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExamUpdate(c *gin.Context) {
	var req ExamUpdateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := ExamUpdateResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateExamUpdateRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	if err := dal.UpdateExam(&model.Exam{
		ID:        int(req.ExamId),
		Name:      req.ExamName,
		BeginTime: time.Unix(req.ExamBeginTime, 0),
		EndTime:   time.Unix(req.ExamEndTime, 0),
		PaperID:   int(req.PaperId),
	}); err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "修改失败："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type ExamUpdateRequest struct {
	ExamId        int32  `json:"exam_id"`
	ExamName      string `json:"exam_name"`
	ExamBeginTime int64  `json:"exam_begin_time"`
	ExamEndTime   int64  `json:"exam_end_time"`
	PaperId       int32  `json:"paper_id"`
}

type ExamUpdateResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateExamUpdateRequest(req *ExamUpdateRequest) error {
	if req.ExamName == "" {
		return fmt.Errorf("名字不能为空")
	}
	if req.PaperId <= 0 {
		return fmt.Errorf("非法的考卷id")
	}
	_, err := dal.GetPaperById(req.PaperId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("考卷ID 不存在")
		}
		return err
	}
	if req.ExamBeginTime <= 0 || req.ExamEndTime <= 0 || req.ExamBeginTime > req.ExamEndTime {
		return fmt.Errorf("考试时间不得为空且考试结束时间不得早于开始时间")
	}
	return nil
}
