package dal

import (
	"bishe/backend/model"
)

func GetStudentByNumber(number string) (*model.Student, error) {
	db := dal.db
	student := model.Student{}
	if err := db.Table(dal.studentTableName).Where("number = ?", number).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
