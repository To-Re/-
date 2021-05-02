package paper

import (
	"bishe/backend/dal"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func PaperDetail(c *gin.Context) {
	var req PaperDetailRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := PaperDetailResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidatePaperDetailRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	paperInfo, err := dal.GetPaperById(req.PaperId)
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询失败："+err.Error()))
		return
	}

	resp.PaperName = paperInfo.Name
	resp.ScoreLimit = int32(paperInfo.ScoreLimit)
	c.JSON(200, resp)
}

type PaperDetailRequest struct {
	PaperId int32 `form:"paper_id"`
}

type PaperDetailResponse struct {
	PaperName  string `json:"paper_name"`
	ScoreLimit int32  `json:"score_limit"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func ValidatePaperDetailRequest(req *PaperDetailRequest) error {
	return nil
}
