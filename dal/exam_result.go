package dal

import (
	"bishe/backend/model"

	"gorm.io/gorm"
)

func GetExamResultList(userId int32, ids []int32) ([]*model.ExamResult, error) {
	db := dal.db
	examResultList := []*model.ExamResult{}
	if err := db.Table(dal.examResultTableName).
		Where("exam_id in (?)", ids).
		Where("student_id = ?", userId).
		Find(&examResultList).Error; err != nil {
		return nil, err
	}
	return examResultList, nil
}

func CreateExamResult(tx *gorm.DB, examResult *model.ExamResult) error {
	if err := tx.Table(dal.examResultTableName).Create(examResult).Error; err != nil {
		return err
	}
	return nil
}

func GetExamResult(req *model.ExamResult) (*model.ExamResult, error) {
	db := dal.db
	res := &model.ExamResult{}
	if err := db.Table(dal.examResultTableName).
		Where("exam_id = ?", req.ExamID).
		Where("student_id = ?", req.StudentID).
		First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetExamResultByStudentIdsExamId(studentIds []int32, exam_id int32) ([]*model.ExamResult, error) {
	db := dal.db
	examResultList := []*model.ExamResult{}
	if err := db.Table(dal.examResultTableName).
		Where("student_id in (?)", studentIds).
		Where("exam_id = (?)", exam_id).
		Find(&examResultList).Error; err != nil {
		return nil, err
	}
	return examResultList, nil
}
