package dal

import (
	"bishe/backend/model"
)

func GetExamKlassListByExamId(exam_id int32) ([]*model.KlassExam, error) {
	db := dal.db
	list := []*model.KlassExam{}
	if err := db.Table(dal.klassExamTableName).Where("exam_id = ?", exam_id).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func CreateKlassExam(klassExam *model.KlassExam) error {
	db := dal.db
	if err := db.Table(dal.klassExamTableName).Create(klassExam).Error; err != nil {
		return err
	}
	return nil
}

func DeleteKlassExam(klassExam *model.KlassExam) error {
	db := dal.db
	if err := db.Table(dal.klassExamTableName).
		Where("klass_id = ?", klassExam.KlassID).
		Where("exam_id = ?", klassExam.ExamID).
		Delete(klassExam).
		Error; err != nil {
		return err
	}
	return nil
}

func GetExamKlassListByKlassId(klass_id int32) ([]*model.KlassExam, error) {
	db := dal.db
	list := []*model.KlassExam{}
	if err := db.Table(dal.klassExamTableName).Where("klass_id = ?", klass_id).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func GetExamKlasstByExamIdKlassId(exam_id, klass_id int32) (*model.KlassExam, error) {
	db := dal.db
	klassExam := model.KlassExam{}
	if err := db.Table(dal.klassExamTableName).
		Where("klass_id = ?", klass_id).
		Where("exam_id = ?", exam_id).
		First(&klassExam).Error; err != nil {
		return nil, err
	}
	return &klassExam, nil
}
