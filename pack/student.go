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

func GetStudentExamList(userId int32) ([]*StudentExam, error) {
	// 拿到班级id
	student, err := dal.GetStudentByUserId(userId)
	if err != nil {
		return nil, err
	}
	examKlassList, err := dal.GetExamKlassListByKlassId(int32(student.KlassID))
	if err != nil {
		return nil, err
	}
	examIds := make([]int32, 0, len(examKlassList))
	for _, v := range examKlassList {
		examIds = append(examIds, int32(v.ExamID))
	}
	// 考试列表
	examInfoList, err := dal.GetExamListByIds(examIds)
	if err != nil {
		return nil, err
	}
	// 得到成绩列表
	scoreInfoList, err := dal.GetExamResultList(userId, examIds)
	if err != nil {
		return nil, err
	}
	scoreMap := make(map[int32]int32)
	for _, v := range scoreInfoList {
		scoreMap[int32(v.ExamID)] = int32(v.Score)
	}

	resList := make([]*StudentExam, 0, len(examInfoList))
	for _, v := range examInfoList {
		tmp := StudentExam{
			ExamId:        int32(v.ID),
			ExamName:      v.Name,
			ExamBeginTime: v.BeginTime.Unix(),
			ExamEndTime:   v.EndTime.Unix(),
		}
		if score, ok := scoreMap[int32(v.ID)]; ok {
			tmp.ExamStudentScore = &score
		}

		resList = append(resList, &tmp)
	}
	return resList, nil
}

type StudentExam struct {
	ExamId           int32
	ExamName         string
	ExamBeginTime    int64
	ExamEndTime      int64
	ExamStudentScore *int32
}
