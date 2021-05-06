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
