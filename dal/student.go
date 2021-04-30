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

func GetStudentByUserId(uid int32) (*model.Student, error) {
	db := dal.db
	student := model.Student{}
	if err := db.Table(dal.studentTableName).Where("id = ?", uid).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func CreateStudent(user *model.Student) error {
	db := dal.db
	if err := db.Table(dal.studentTableName).Create(user).Error; err != nil {
		return err
	}
	return nil
}
