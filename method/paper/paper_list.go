package paper

import (
	"bishe/backend/dal"
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func PaperList(c *gin.Context) {
	var req PaperListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := PaperListResponse{
		Code: 0,
		Msg:  "",
	}
	if err := ValidatePaperListRequest(&req); err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]+": "+err.Error()))
		return
	}

	papers, err := dal.GetPaperList()
	if err != nil {
		c.JSON(200, util.BuildError(util.FUNCFAILURE, "查询考卷列表失败："+err.Error()))
		return
	}

	resPapers := make([]*Paper, 0, len(papers))

	for _, v := range papers {
		resPapers = append(resPapers, &Paper{
			PaperId:    int32(v.ID),
			PaperName:  v.Name,
			ScoreLimit: int32(v.ScoreLimit),
		})
	}
	resp.Papers = resPapers
	c.JSON(200, resp)
}

type PaperListRequest struct {
}

type PaperListResponse struct {
	Papers []*Paper `json:"papers"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type Paper struct {
	PaperId    int32  `json:"paper_id"`
	PaperName  string `json:"paper_name"`
	ScoreLimit int32  `json:"score_limit"`
}

func ValidatePaperListRequest(req *PaperListRequest) error {
	return nil
}
