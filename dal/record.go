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
