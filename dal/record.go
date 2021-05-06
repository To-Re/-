package dal

import (
	"bishe/backend/model"

	"gorm.io/gorm"
)

func CreateRecord(tx *gorm.DB, records *[]model.Record) error {
	if err := tx.Table(dal.recordTableName).Create(records).Error; err != nil {
		return err
	}
	return nil
}

func GetRecordList(req *model.Record) ([]*model.Record, error) {
	db := dal.db
	list := []*model.Record{}
	if err := db.Table(dal.recordTableName).
		Where("student_id = ?", req.StudentID).
		Where("exam_id = ?", req.ExamID).
		Where("paper_id = ?", req.PaperID).
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
