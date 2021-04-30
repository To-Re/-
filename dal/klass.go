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
