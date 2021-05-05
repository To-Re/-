package pack

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"fmt"

	"gorm.io/gorm"
)

func ExamKlassList(examId int32) ([]*model.Klass, error) {
	examKlassList, err := dal.GetExamKlassListByExamId(examId)
	if err != nil {
		return nil, err
	}
	klassIds := make([]int32, 0, len(examKlassList))
	for _, v := range examKlassList {
		klassIds = append(klassIds, int32(v.KlassID))
	}

	return dal.GetExamListByIds(klassIds)
}

func ExamKlassBind(req *model.KlassExam) error {
	_, err := dal.GetKlassById(int32(req.KlassID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("班级id 不存在")
		}
		return err
	}
	return dal.CreateKlassExam(req)
}
