package pack

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"fmt"

	"gorm.io/gorm"
)

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

func UpdateStudentInfo(req *StructInfoUpdate) error {
	_, err := dal.GetKlassById(req.KlassId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("班级id 不存在")
		}
		return err
	}
	return dal.UpdateStudent(&model.Student{
		ID:       int(req.StudentId),
		Name:     req.StudentName,
		Password: req.Password,
		KlassID:  int(req.KlassId),
	})
}

type StructInfoUpdate struct {
	StudentId   int32
	StudentName string
	KlassId     int32
	Password    string
}
