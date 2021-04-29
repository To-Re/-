package dal

import (
	"bishe/backend/model"
)

func GetKlassList() ([]*model.Klass, error) {
	db := dal.db
	klassList := []*model.Klass{}
	if err := db.Table(dal.klassTableName).Find(&klassList).Error; err != nil {
		return nil, err
	}
	return klassList, nil
}
