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
