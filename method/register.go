package method

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var req RegisterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := RegisterResponse{
		Code: 0,
		Msg:  "",
	}
	if req.UserType == util.UserTypeTeacher {
		user, err := dal.GetTeacherByNumber(req.Number)
		if err != gorm.ErrRecordNotFound && err != nil {
			c.JSON(200, util.BuildError(util.NETWORKERROR, util.ErrMap[util.NETWORKERROR]+": "+err.Error()))
			return
		}
		if user != nil {
			c.JSON(200, util.BuildError(util.REGISTERERROR, util.ErrMap[util.REGISTERERROR]+"：number 已存在"))
			return
		}

		if err := dal.CreateTeacher(&model.Teacher{
			Name:     req.Name,
			Number:   req.Number,
			Password: req.Password,
		}); err != nil {
			c.JSON(200, util.BuildError(util.NETWORKERROR, util.ErrMap[util.NETWORKERROR]+": "+err.Error()))
			return
		}
		fmt.Println(user)
	} else if req.UserType == util.UserTypeStudent {
	} else {
		c.JSON(200, util.BuildError(util.REGISTERERROR, util.ErrMap[util.REGISTERERROR]+"：请检查人员类型"))
		return
	}
	c.JSON(200, resp)
}

type RegisterRequest struct {
	Number   string `json:"number"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
	Name     string `json:"name"`
}
type RegisterResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
