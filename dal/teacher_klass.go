package dal

import (
	"bishe/backend/model"

	"gorm.io/gorm"
)

func GetTeacherKlassListByKlassIds(klassIds []int32) ([]*model.TeacherKlass, error) {
	db := dal.db
	teacherKlassList := []*model.TeacherKlass{}
	if err := db.Table(dal.teacherKlassTableName).Where("klass_id in (?)", klassIds).Find(&teacherKlassList).Error; err != nil {
		return nil, err
	}
	return teacherKlassList, nil
}

func CreateTeacherKlass(tx *gorm.DB, teacherKlass *model.TeacherKlass) error {
	if err := tx.Table(dal.teacherKlassTableName).Create(teacherKlass).Error; err != nil {
		return err
	}
	return nil
}
