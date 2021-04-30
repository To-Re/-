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

func GetTeacherByUserId(uid int32) (*model.Teacher, error) {
	db := dal.db
	teacher := model.Teacher{}
	if err := db.Table(dal.teacherTableName).Where("id = ?", uid).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func GetTeacherListByIds(Ids []int32) ([]*model.Teacher, error) {
	db := dal.db
	teacherList := []*model.Teacher{}
	if err := db.Table(dal.teacherTableName).Where("id in (?)", Ids).Find(&teacherList).Error; err != nil {
		return nil, err
	}
	return teacherList, nil
}

func GetTeacherMapByIds(Ids []int32) (map[int32]*model.Teacher, error) {
	list, err := GetTeacherListByIds(Ids)
	if err != nil {
		return nil, err
	}
	teacherMap := make(map[int32]*model.Teacher)
	for _, v := range list {
		teacherMap[int32(v.ID)] = v
	}
	return teacherMap, nil
}

func CreateTeacher(user *model.Teacher) error {
	db := dal.db
	if err := db.Table(dal.teacherTableName).Create(user).Error; err != nil {
		return err
	}
	return nil
}
