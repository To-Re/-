package pack

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"time"
)

func GetResultList() ([]*Result, error) {
	examListInfo, err := dal.GetExamList(time.Now().Unix())
	if err != nil {
		return nil, err
	}
	examIds := make([]int32, 0, len(examListInfo))
	examMap := make(map[int32]*model.Exam)
	for _, v := range examListInfo {
		examIds = append(examIds, int32(v.ID))
		examMap[int32(v.ID)] = v
	}

	klassList, err := dal.GetExamKlassListByExamIds(examIds)
	if err != nil {
		return nil, err
	}
	klassIds := make([]int32, 0, len(examListInfo))
	klassExamMap := make(map[int32]bool)
	for _, v := range klassList {
		if _, ok := klassExamMap[int32(v.KlassID)]; !ok {
			klassExamMap[int32(v.KlassID)] = true
			klassIds = append(klassIds, int32(v.KlassID))
		}
	}
	klassInfoList, err := dal.GetKlassListByIds(klassIds)
	klassInfoMap := make(map[int32]*model.Klass)
	for _, v := range klassInfoList {
		klassInfoMap[int32(v.ID)] = v
	}

	resp := make([]*Result, 0)
	for _, v := range klassList {
		tmp := examMap[int32(v.ExamID)]
		resp = append(resp, &Result{
			KlassId:       int32(v.KlassID),
			ExamId:        int32(v.ExamID),
			ExamName:      tmp.Name,
			ExamBeginTime: tmp.BeginTime.Unix(),
			ExamEndTime:   tmp.EndTime.Unix(),
			KlassName:     klassInfoMap[int32(v.KlassID)].Name,
		})
	}
	return resp, nil
}

type Result struct {
	ExamId        int32  `json:"exam_id"`
	ExamName      string `json:"exam_name"`
	ExamBeginTime int64  `json:"exam_begin_time"`
	ExamEndTime   int64  `json:"exam_end_time"`
	KlassId       int32  `json:"klass_id"`
	KlassName     string `json:"klass_name"`
}

func StudentResultList(exam_id, klass_id int32) ([]*StudentResult, error) {
	// 获取学生info list
	studentInfo, err := dal.GetStudentByKlassId(klass_id)
	if err != nil {
		return nil, err
	}
	studentIds := make([]int32, 0, len(studentInfo))
	studentMap := make(map[int32]*model.Student)
	for _, v := range studentInfo {
		studentIds = append(studentIds, int32(v.ID))
		studentMap[int32(v.ID)] = v
	}
	// exam_result 获取分数和考试状态
	// 最好在限制 paper_id 懒得弄了
	examResultInfo, err := dal.GetExamResultByStudentIdsExamId(studentIds, exam_id)
	if err != nil {
		return nil, err
	}
	examResultMap := make(map[int32]*model.ExamResult)
	for _, v := range examResultInfo {
		examResultMap[int32(v.StudentID)] = v
	}

	resp := make([]*StudentResult, 0, len(studentIds))
	for _, v := range studentIds {
		tmp := &StudentResult{
			StudentId:   v,
			StudentName: studentMap[v].Name,
		}

		if e, ok := examResultMap[v]; ok {
			tmp.StudentResultStatus = "已交卷"
			tmp.StudentScore = int32(e.Score)
		} else {
			tmp.StudentResultStatus = "未考试"
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

type StudentResult struct {
	StudentId           int32  `json:"student_id"`
	StudentName         string `json:"student_name"`
	StudentResultStatus string `json:"student_result_status"`
	StudentScore        int32  `json:"student_score"`
}
