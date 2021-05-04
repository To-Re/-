package student

import (
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func StudentInfo(c *gin.Context) {
	var req StudentInfoRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := StudentInfoResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateStudentInfoRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	userId, err := util.GetId(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
		return
	}

	studentInfo, err := pack.GetStudentInfoById(userId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询失败："+err.Error()))
		return
	}

	resp.StudentId = studentInfo.StudentId
	resp.StudentName = studentInfo.StudentName
	resp.StudentNumber = studentInfo.StudentNumber
	resp.KlassId = studentInfo.KlassId
	resp.KlassName = studentInfo.KlassName
	c.JSON(200, resp)
}

type StudentInfoRequest struct {
}

type StudentInfoResponse struct {
	StudentId     int32  `json:"student_id"`
	StudentName   string `json:"student_name"`
	StudentNumber string `json:"student_number"`
	KlassId       int32  `json:"klass_id"`
	KlassName     string `json:"klass_name"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateStudentInfoRequest(req *StudentInfoRequest) error {
	return nil
}
