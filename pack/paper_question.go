package pack

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
	"fmt"

	"gorm.io/gorm"
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

func PaperQuestionBind(req *model.PaperQuestion) error {
	// 开启事务
	tx := dal.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// 1.验证试卷题目绑定关系
	_, err := dal.GetPaperQuestionByPaperIdQuestionId(int32(req.PaperID), int32(req.QuestionID))
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}
	// 2.拿到 paper 表中试卷分数
	paperInfo, err := dal.GetPaperById(int32(req.PaperID))
	if err != nil {
		tx.Rollback()
		return err
	}
	// 3.检查是否有该题目
	_, err = dal.GetQuestionById(int32(req.QuestionID))
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("题目id 不存在")
		}
		return err
	}
	// 4.修改 paper 表中试卷分数
	paperInfo.ScoreLimit += req.QuestionScore
	err = dal.UpdatePaperScore(tx, paperInfo)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 5.创建试卷题目绑定关系
	err = dal.CreatePaperQuestion(tx, req)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func PaperQuestionDelete(req *model.PaperQuestion) error {
	// 开启事务
	tx := dal.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// 1.查询考卷信息
	paperInfo, err := dal.GetPaperById(int32(req.PaperID))
	if err != nil {
		tx.Rollback()
		return err
	}
	// 2.查询该条考卷-题目信息
	paperQuestionInfo, err := dal.GetPaperQuestionByPaperIdQuestionId(int32(req.PaperID), int32(req.QuestionID))
	if err != nil {
		tx.Rollback()
		return err
	}
	// 3.更新考卷得分
	paperInfo.ScoreLimit -= paperQuestionInfo.QuestionScore
	err = dal.UpdatePaperScore(tx, paperInfo)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 4.删除关系
	err = dal.DeletePaperQuestionByPaperIdQuestionId(tx, paperQuestionInfo)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
