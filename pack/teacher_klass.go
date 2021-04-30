package pack

import (
	"bishe/backend/dal"
	"bishe/backend/model"
)

func GetTeacherModelMapbyKlassIds(klassIds []int32) (map[int32]*model.Teacher, error) {
	// teacher-klass list
	teacherKlassList, err := dal.GetTeacherKlassListByKlassIds(klassIds)
	if err != nil {
		return nil, err
	}
	teacherIdList := make([]int32, 0, len(klassIds))
	for _, v := range teacherKlassList {
		teacherIdList = append(teacherIdList, int32(v.TeacherID))
	}

	// key teacher_id value teacher_model
	teacherMapKeyIsTeacherId, err := dal.GetTeacherMapByIds(teacherIdList)
	if err != nil {
		return nil, err
	}

	// key klass_id value teacher_model
	teacherMapKeyIsKlassId := make(map[int32]*model.Teacher)
	for _, v := range teacherKlassList {
		teacherMapKeyIsKlassId[int32(v.KlassID)] = teacherMapKeyIsTeacherId[int32(v.TeacherID)]
	}
	return teacherMapKeyIsKlassId, nil
}

func CreateKlassAndBindTeacher(userId int32, klassName string) error {
	// 开启事物
	tx := dal.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	klass := &model.Klass{
		Name: klassName,
	}
	if err := dal.CreateKlass(tx, klass); err != nil {
		tx.Rollback()
		return err
	}

	if err := dal.CreateTeacherKlass(tx, &model.TeacherKlass{
		TeacherID: int(userId),
		KlassID:   klass.ID,
	}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
