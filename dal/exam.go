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

func CreateExam(exam *model.Exam) error {
	db := dal.db
	if err := db.Table(dal.examTableName).Create(exam).Error; err != nil {
		return err
	}
	return nil
}

func GetExamById(exam_id int32) (*model.Exam, error) {
	db := dal.db
	exam := &model.Exam{}
	if err := db.Table(dal.examTableName).Where("id = ?", exam_id).First(&exam).Error; err != nil {
		return nil, err
	}
	return exam, nil
}

func UpdateExam(exam *model.Exam) error {
	db := dal.db
	if err := db.Table(dal.examTableName).
		Where("id = ?", exam.ID).
		Updates(exam).Error; err != nil {
		return err
	}
	return nil
}
