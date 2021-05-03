package dal

import (
	"bishe/backend/model"

	"gorm.io/gorm"
)

func GetPaperQuestionListByPaperId(paperId int32) ([]*model.PaperQuestion, error) {
	db := dal.db
	list := []*model.PaperQuestion{}
	if err := db.Table(dal.paperQuestionTableName).Where("paper_id = ?", paperId).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func GetPaperQuestionListByPaperIdQuestionId(paperId, questionId int32) (*model.PaperQuestion, error) {
	db := dal.db
	res := model.PaperQuestion{}
	if err := db.Table(dal.paperQuestionTableName).
		Where("paper_id = ?", paperId).
		Where("question_id = ?", questionId).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func CreatePaperQuestion(tx *gorm.DB, paperQuestion *model.PaperQuestion) error {
	if err := tx.Table(dal.paperQuestionTableName).Create(paperQuestion).Error; err != nil {
		return err
	}
	return nil
}
