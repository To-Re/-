package dal

import "bishe/backend/model"

func GetQuestionList() ([]*model.Question, error) {
	db := dal.db
	list := []*model.Question{}
	if err := db.Table(dal.questionTableName).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
