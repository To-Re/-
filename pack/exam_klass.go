package pack

import (
	"bishe/backend/dal"
	"bishe/backend/model"
)

func ExamKlassList(examId int32) ([]*model.Klass, error) {
	examKlassList, err := dal.GetExamKlassListByKlassId(examId)
	if err != nil {
		return nil, err
	}
	klassIds := make([]int32, 0, len(examKlassList))
	for _, v := range examKlassList {
		klassIds = append(klassIds, int32(v.KlassID))
	}

	return dal.GetExamListByIds(klassIds)
}
