package dal

import "bishe/backend/model"

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
