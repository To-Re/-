package pack

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
)

func GetPaperQuestionList(paperId int32) ([]*PaperQuestion, error) {
	// 获取 paper_question 基础信息
	paperQusetions, err := dal.GetPaperQuestionListByPaperId(paperId)
	if err != nil {
		return nil, err
	}
	questionIds := make([]int32, 0, len(paperQusetions))
	res := make([]*PaperQuestion, 0, len(paperQusetions))
	for _, v := range paperQusetions {
		questionIds = append(questionIds, int32(v.QuestionID))
		res = append(res, &PaperQuestion{
			QuestionId:    int32(v.QuestionID),
			QuestionScore: int32(v.QuestionScore),
		})
	}

	// 读取 question 表组装map
	Questions, err := dal.GetQuestionListByIds(questionIds)
	if err != nil {
		return nil, err
	}
	questionMap := make(map[int32]*model.Question)
	for _, v := range Questions {
		questionMap[int32(v.ID)] = v
	}

	// 拿到 question 信息
	for _, v := range res {
		questionInfo := questionMap[v.QuestionId]
		v.QuestionDesc = questionInfo.Desc
		v.QuestionType = util.QuestionTypeMap[int32(questionInfo.Type)]
	}
	return res, nil
}

type PaperQuestion struct {
	QuestionId    int32
	QuestionScore int32
	QuestionDesc  string
	QuestionType  string
}
