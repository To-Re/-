package dal

import "bishe/backend/model"

func GetExamList() ([]*model.Exam, error) {
	db := dal.db
	list := []*model.Exam{}
	if err := db.Table(dal.examTableName).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
