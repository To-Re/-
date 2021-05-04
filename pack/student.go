package pack

import "bishe/backend/dal"

func GetStudentInfoById(student_id int32) (*StructInfo, error) {
	student, err := dal.GetStudentByUserId(student_id)
	if err != nil {
		return nil, err
	}
	klass, err := dal.GetKlassById(int32(student.KlassID))
	if err != nil {
		return nil, err
	}
	return &StructInfo{
		StudentId:     int32(student.ID),
		StudentName:   student.Name,
		StudentNumber: student.Number,
		KlassId:       int32(student.KlassID),
		KlassName:     klass.Name,
	}, nil
}

type StructInfo struct {
	StudentId     int32
	StudentName   string
	StudentNumber string
	KlassId       int32
	KlassName     string
}
