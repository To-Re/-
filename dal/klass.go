package dal

import (
	"bishe/backend/model"

	"gorm.io/gorm"
)

func GetKlassList() ([]*model.Klass, error) {
	db := dal.db
	klassList := []*model.Klass{}
	if err := db.Table(dal.klassTableName).Find(&klassList).Error; err != nil {
		return nil, err
	}
	return klassList, nil
}

func CreateKlass(tx *gorm.DB, klass *model.Klass) error {
	if err := tx.Table(dal.klassTableName).Create(klass).Error; err != nil {
		return err
	}
	return nil
}

func GetKlassById(klass_id int32) (*model.Klass, error) {
	db := dal.db
	klass := &model.Klass{}
	if err := db.Table(dal.klassTableName).Where("id = ?", klass_id).First(&klass).Error; err != nil {
		return nil, err
	}
	return klass, nil
}

func UpdateKlass(klass *model.Klass) error {
	db := dal.db
	if err := db.Table(dal.klassTableName).
		Where("id = ?", klass.ID).
		Updates(klass).Error; err != nil {
		return err
	}
	return nil
}

func GetExamListByIds(ids []int32) ([]*model.Klass, error) {
	db := dal.db
	klassList := []*model.Klass{}
	if err := db.Table(dal.klassTableName).Where("id in (?)", ids).Find(&klassList).Error; err != nil {
		return nil, err
	}
	return klassList, nil
}
