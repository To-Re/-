package dal

import (
	"bishe/backend/model"
)

func GetQuestionList() ([]*model.Question, error) {
	db := dal.db
	list := []*model.Question{}
	if err := db.Table(dal.questionTableName).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func CreateQuestion(question *model.Question) error {
	db := dal.db
	if err := db.Table(dal.questionTableName).Create(question).Error; err != nil {
		return err
	}
	return nil
}

func GetQuestionById(question_id int32) (*model.Question, error) {
	db := dal.db
	question := &model.Question{}
	if err := db.Table(dal.questionTableName).Where("id = ?", question_id).Find(&question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func UpdateQuestion(question *model.Question) error {
	db := dal.db
	if err := db.Table(dal.questionTableName).
		Where("id = ?", question.ID).
		Updates(question).Error; err != nil {
		return err
	}
	return nil
}
