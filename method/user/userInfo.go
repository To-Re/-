package user

import (
	"bishe/backend/dal"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var req UserInfoRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := UserInfoResponse{
		Code: 0,
		Msg:  "",
	}
	userId, err := util.GetId(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
		return
	}
	userType, err := util.GetWho(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
		return
	}
	userName := ""
	if userType == util.UserTypeTeacher {
		teacher, err := dal.GetTeacherByUserId(userId)
		if err != nil || teacher == nil {
			c.JSON(200, util.BuildError(util.LOGINERROR, util.ErrMap[util.LOGINERROR]+"：账号不存在"))
			return
		}
		userName = teacher.Name
	} else if userType == util.UserTypeStudent {
		student, err := dal.GetStudentByUserId(userId)
		if err != nil || student == nil {
			c.JSON(200, util.BuildError(util.LOGINERROR, util.ErrMap[util.LOGINERROR]+"：账号不存在"))
			return
		}
		userName = student.Name
	} else {
		c.JSON(200, util.BuildError(util.LOGINERROR, util.ErrMap[util.LOGINERROR]+"：请检查人员类型"))
		return
	}

	resp.UserId = userId
	resp.UserType = userType
	resp.Name = userName
	c.JSON(200, resp)
}

type UserInfoRequest struct {
}

type UserInfoResponse struct {
	Name     string `json:"name"`
	UserId   int32  `json:"user_id"`
	UserType int32  `json:"user_type"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
