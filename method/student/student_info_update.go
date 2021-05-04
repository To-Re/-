package student

import (
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func StudentInfoUpdate(c *gin.Context) {
	var req StudentInfoUpdateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := StudentInfoUpdateResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidateStudentInfoUpdateRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	userId, err := util.GetId(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
		return
	}

	if err := pack.UpdateStudentInfo(&pack.StructInfoUpdate{
		StudentId:   userId,
		StudentName: req.StudentName,
		KlassId:     req.KlassId,
		Password:    req.Password,
	}); err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "修改失败："+err.Error()))
		return
	}

	c.JSON(200, resp)
}

type StudentInfoUpdateRequest struct {
	StudentName string `json:"student_name"`
	KlassId     int32  `json:"klass_id"`
	Password    string `json:"password"`
}

type StudentInfoUpdateResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidateStudentInfoUpdateRequest(req *StudentInfoUpdateRequest) error {
	return nil
}
