package exam

import (
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func ExamKlassList(c *gin.Context) {
	var req ExamKlassListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := ExamKlassListResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateExamKlassListRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	examKlasses, err := pack.ExamKlassList(req.ExamId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询考试班级列表失败："+err.Error()))
		return
	}

	res := make([]*Klass, 0, len(examKlasses))
	for _, v := range examKlasses {
		res = append(res, &Klass{
			KlassId:   int32(v.ID),
			KlassName: v.Name,
		})
	}
	resp.Klasses = res

	c.JSON(200, resp)
}

type ExamKlassListRequest struct {
	ExamId int32 `form:"exam_id"`
}

type ExamKlassListResponse struct {
	Klasses []*Klass `json:"klasses"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type Klass struct {
	KlassId   int32  `json:"klass_id"`
	KlassName string `json:"klass_name"`
}

func ValidateExamKlassListRequest(req *ExamKlassListRequest) error {
	return nil
}
