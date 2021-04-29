package dal

import (
	"bishe/backend/model"
)

func GetTeacherByNumber(number string) (*model.Teacher, error) {
	db := dal.db
	teacher := model.Teacher{}
	if err := db.Table(dal.teacherTableName).Where("number = ?", number).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}