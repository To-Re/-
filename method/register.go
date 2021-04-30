package method

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"

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

	if req.Name == "" || req.Number == "" || req.Password == "" {
		c.JSON(200, util.BuildError(util.REGISTERERROR, util.ErrMap[util.REGISTERERROR]+": 账号/密码/姓名 不得为空"))
		return
	}

	if err := register(&req); err != nil {
		c.JSON(200, err)
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

func register(req *RegisterRequest) *util.ErrorInfo {
	if req.UserType == util.UserTypeTeacher {
		user, err := dal.GetTeacherByNumber(req.Number)
		if err != gorm.ErrRecordNotFound && err != nil {
			return util.BuildError(util.NETWORKERROR, util.ErrMap[util.NETWORKERROR]+": "+err.Error())
		}
		if user != nil {
			return util.BuildError(util.REGISTERERROR, util.ErrMap[util.REGISTERERROR]+"：number 已存在")
		}
		if err := dal.CreateTeacher(&model.Teacher{
			Name:     req.Name,
			Number:   req.Number,
			Password: req.Password,
		}); err != nil {
			return util.BuildError(util.NETWORKERROR, util.ErrMap[util.NETWORKERROR]+": "+err.Error())
		}
	} else if req.UserType == util.UserTypeStudent {
		user, err := dal.GetStudentByNumber(req.Number)
		if err != gorm.ErrRecordNotFound && err != nil {
			return util.BuildError(util.NETWORKERROR, util.ErrMap[util.NETWORKERROR]+": "+err.Error())
		}
		if user != nil {
			return util.BuildError(util.REGISTERERROR, util.ErrMap[util.REGISTERERROR]+"：number 已存在")
		}
		if err := dal.CreateStudent(&model.Student{
			Name:     req.Name,
			Number:   req.Number,
			Password: req.Password,
			KlassID:  0,
		}); err != nil {
			return util.BuildError(util.NETWORKERROR, util.ErrMap[util.NETWORKERROR]+": "+err.Error())
		}
	} else {
		return util.BuildError(util.REGISTERERROR, util.ErrMap[util.REGISTERERROR]+"：请检查人员类型")
	}
	return nil
}
