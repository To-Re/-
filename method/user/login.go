package user

import (
	"bishe/backend/dal"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := LoginResponse{
		Code: 0,
		Msg:  "",
	}
	// 校验
	var userId int = 0
	if req.UserType == util.UserTypeTeacher { // 老师
		teacher, err := dal.GetTeacherByNumber(req.Number)
		if err != nil || teacher == nil || teacher.ID <= 0 || req.Password != teacher.Password {
			c.JSON(200, util.BuildError(util.LOGINERROR, util.ErrMap[util.LOGINERROR]+"：请检查账号或密码"))
			return
		}
		userId = teacher.ID
	} else if req.UserType == util.UserTypeStudent { // 学生
		student, err := dal.GetStudentByNumber(req.Number)
		if err != nil || student == nil || student.ID <= 0 || req.Password != student.Password {
			c.JSON(200, util.BuildError(util.LOGINERROR, util.ErrMap[util.LOGINERROR]+"：请检查账号或密码"))
			return
		}
		userId = student.ID
	} else {
		c.JSON(200, util.BuildError(util.LOGINERROR, util.ErrMap[util.LOGINERROR]+"：请检查人员类型"))
		return
	}
	ret, err := util.CreteToken(int32(userId), int32(req.UserType))
	if err != nil {
		// 报警
		c.JSON(200, util.BuildError(util.LOGINERROR, util.ErrMap[util.LOGINERROR]+"：出现带问题，请联系管理员"))
		return
	}
	resp.Token = ret
	fmt.Println("token : " + ret)
	c.JSON(200, resp)
}

type LoginRequest struct {
	Number   string `json:"number"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
}
type LoginResponse struct {
	Token string `json:"token"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
