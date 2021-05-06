package dal

import (
	"bishe/backend/model"
	"time"
)

func GetExamList(endTime int64) ([]*model.Exam, error) {
	db := dal.db
	list := []*model.Exam{}
	db = db.Table(dal.examTableName)
	if endTime > 0 {
		db = db.Where("end_time < ?", time.Unix(endTime, 0))
	}
	if err := db.Find(&list).Error; err != nil {
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

func GetExamListByIds(ids []int32) ([]*model.Exam, error) {
	db := dal.db
	examList := []*model.Exam{}
	if err := db.Table(dal.examTableName).Where("id in (?)", ids).Find(&examList).Error; err != nil {
		return nil, err
	}
	return examList, nil
}
