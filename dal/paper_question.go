package dal

import "bishe/backend/model"

func GetPaperQuestionListByPaperId(paperId int32) ([]*model.PaperQuestion, error) {
	db := dal.db
	list := []*model.PaperQuestion{}
	if err := db.Table(dal.paperQuestionTableName).Where("paper_id = ?", paperId).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
