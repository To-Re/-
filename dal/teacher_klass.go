package dal

import "bishe/backend/model"

func GetTeacherKlassListByKlassIds(klassIds []int32) ([]*model.TeacherKlass, error) {
	db := dal.db
	teacherKlassList := []*model.TeacherKlass{}
	if err := db.Table(dal.teacherKlassTableName).Where("klass_id in (?)", klassIds).Find(&teacherKlassList).Error; err != nil {
		return nil, err
	}
	return teacherKlassList, nil
}
