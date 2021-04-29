package method

import (
	"bishe/backend/dal"
	"bishe/backend/pack"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func KlassList(c *gin.Context) {
	var req KlassListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := KlassListResponse{
		Code: 0,
		Msg:  "",
	}
	klassList, err := dal.GetKlassList()
	if err != nil {
		c.JSON(200, util.BuildError(util.NETWORKERROR, util.ErrMap[util.NETWORKERROR]))
		return
	}

	klasses := make([]*Klasses, 0, len(klassList))
	klassIdList := make([]int32, 0, len(klassList))

	for _, v := range klassList {
		klasses = append(klasses, &Klasses{
			KlassId:   int32(v.ID),
			KlassName: v.Name,
		})
		klassIdList = append(klassIdList, int32(v.ID))
	}
	teacherMapKeyIsKlassId, err := pack.GetTeacherModelMapbyKlassIds(klassIdList)
	if err != nil {
		c.JSON(200, util.BuildError(util.NETWORKERROR, util.ErrMap[util.NETWORKERROR]))
		return
	}
	for _, v := range klasses {
		teacher := teacherMapKeyIsKlassId[v.KlassId]
		if teacher == nil {
			continue
		}
		v.TeacherName = teacher.Name
	}

	resp.Klasses = klasses
	c.JSON(200, resp)
}

type KlassListRequest struct {
}

type KlassListResponse struct {
	Total   int32      `json:"total"`
	Page    int32      `json:"page"`
	Size    int32      `json:"size"`
	Klasses []*Klasses `json:"klasses"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type Klasses struct {
	KlassId     int32  `json:"klass_id"`
	KlassName   string `json:"klass_name"`
	TeacherName string `json:"teacher_name"`
}
