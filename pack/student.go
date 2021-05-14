package pack

import (
	"bishe/backend/dal"
	"bishe/backend/model"
	"bishe/backend/util"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GetStudentInfoById(student_id int32) (*StructInfo, error) {
	student, err := dal.GetStudentByUserId(student_id)
	if err != nil {
		return nil, err
	}
	resp := StructInfo{
		StudentId:     int32(student.ID),
		StudentName:   student.Name,
		StudentNumber: student.Number,
	}
	klass, err := dal.GetKlassById(int32(student.KlassID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &resp, nil
		}
		return nil, err
	}
	resp.KlassId = util.ConvertInt32ToPtr(int32(student.KlassID))
	resp.KlassName = &klass.Name
	return &resp, nil
}

type StructInfo struct {
	StudentId     int32
	StudentName   string
	StudentNumber string
	KlassId       *int32
	KlassName     *string
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

func GetStudentExamDetail(userId, ExamId int32) (*StudentExamDetail, error) {
	examInfo, err := dal.GetExamById(ExamId)
	if err != nil {
		return nil, err
	}
	// 校验是否开始考试
	if examInfo.BeginTime.Unix() > time.Now().Unix() {
		return nil, fmt.Errorf("考试未开始")
	}
	// 校验该学生是否拥有该考试
	studentInfo, err := dal.GetStudentByUserId(userId)
	if err != nil {
		return nil, err
	}
	_, err = dal.GetExamKlasstByExamIdKlassId(ExamId, int32(studentInfo.KlassID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("没有考试资格")
		}
		return nil, err
	}
	// 获取考卷信息
	paperInfo, err := dal.GetPaperById(int32(examInfo.PaperID))
	if err != nil {
		return nil, err
	}
	// 获取考题ids
	questions, err := dal.GetPaperQuestionListByPaperId(int32(paperInfo.ID))
	if err != nil {
		return nil, err
	}
	questionIds := make([]int32, 0, len(questions))
	questionScoreMap := make(map[int32]int32)
	for _, v := range questions {
		questionIds = append(questionIds, int32(v.QuestionID))
		questionScoreMap[int32(v.QuestionID)] = int32(v.QuestionScore)
	}

	// 获取考题信息
	questionInfo, err := dal.GetQuestionListByIds(questionIds)
	if err != nil {
		return nil, err
	}

	resQuestions := make([]*Question, 0, len(questionIds))
	for _, v := range questionInfo {
		resQuestions = append(resQuestions, &Question{
			QuestionId:    int32(v.ID),
			QuestionDesc:  v.Desc,
			QuestionType:  util.QuestionTypeMap[int32(v.Type)],
			QuestionScore: questionScoreMap[int32(v.ID)],
			OptionDescA:   v.OptionDescA,
			OptionDescB:   v.OptionDescB,
			OptionDescC:   v.OptionDescC,
			OptionDescD:   v.OptionDescD,
		})
	}
	// 组装返回
	resp := &StudentExamDetail{
		ExamName:     examInfo.Name,
		ExamEndTime:  examInfo.EndTime.Unix(),
		ScoreLimit:   int32(paperInfo.ScoreLimit),
		StudentScore: nil,
		Questions:    resQuestions,
	}

	// 尝试获取考试结果
	examResutlInfo, err := dal.GetExamResult(&model.ExamResult{
		StudentID: int(userId),
		ExamID:    int(ExamId),
	})
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		// 没有考试记录
		return resp, nil
	}
	// 有考试记录，尝试获取作答结果
	records, err := dal.GetRecordList(&model.Record{
		ExamID:    int(ExamId),
		StudentID: int(userId),
		PaperID:   paperInfo.ID,
	})
	recordMap := make(map[int32]*model.Record)
	for _, v := range records {
		recordMap[int32(v.QuestionID)] = v
	}
	resp.StudentScore = util.ConvertInt32ToPtr(int32(examResutlInfo.Score))
	for _, v := range resp.Questions {
		v.QuestionStudentScore = util.ConvertInt32ToPtr(int32(recordMap[v.QuestionId].Score))
		v.StudentAnswer = recordMap[v.QuestionId].Desc
	}
	return resp, nil
}

type StudentExamDetail struct {
	ExamName     string      `json:"exam_name"`
	ScoreLimit   int32       `json:"score_limit"`
	ExamEndTime  int64       `json:"exam_end_time"`
	StudentScore *int32      `json:"student_score,omitempty"`
	Questions    []*Question `json:"questions"`
}

type Question struct {
	QuestionId           int32  `json:"question_id"`
	QuestionDesc         string `json:"question_desc"`
	QuestionScore        int32  `json:"question_score"`
	QuestionType         string `json:"question_type"`
	OptionDescA          string `json:"option_desc_a"`
	OptionDescB          string `json:"option_desc_b"`
	OptionDescC          string `json:"option_desc_c"`
	OptionDescD          string `json:"option_desc_d"`
	StudentAnswer        string `json:"student_answer"`
	QuestionStudentScore *int32 `json:"question_student_score,omitempty"`
}

func StudentExamCommit(userId, ExamId int32, studentAnswer []*StudentQeustionAnswer) error {
	examInfo, err := dal.GetExamById(ExamId)
	if err != nil {
		return err
	}
	// 校验考试时间
	if examInfo.BeginTime.Unix() > time.Now().Unix() {
		return fmt.Errorf("考试未开始")
	}
	if examInfo.EndTime.Unix() < time.Now().Unix() {
		return fmt.Errorf("考试结束")
	}
	// 校验考试资格
	studentInfo, err := dal.GetStudentByUserId(userId)
	if err != nil {
		return err
	}
	_, err = dal.GetExamKlasstByExamIdKlassId(ExamId, int32(studentInfo.KlassID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("没有考试资格")
		}
		return err
	}
	// 是否已交卷
	_, err = dal.GetExamResult(&model.ExamResult{
		StudentID: int(userId),
		ExamID:    int(ExamId),
	})
	if err == nil {
		return fmt.Errorf("已交卷")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// 获取考卷信息
	paperInfo, err := dal.GetPaperById(int32(examInfo.PaperID))
	if err != nil {
		return err
	}
	// 获取考题ids
	questions, err := dal.GetPaperQuestionListByPaperId(int32(paperInfo.ID))
	if err != nil {
		return err
	}
	questionIds := make([]int32, 0, len(questions))
	questionScoreMap := make(map[int32]int32)
	for _, v := range questions {
		questionIds = append(questionIds, int32(v.QuestionID))
		questionScoreMap[int32(v.QuestionID)] = int32(v.QuestionScore)
	}
	// 获取考题信息
	questionInfo, err := dal.GetQuestionListByIds(questionIds)
	if err != nil {
		return err
	}
	// 组装考题map
	answerMap := make(map[int32]*model.Question)
	for _, v := range questionInfo {
		answerMap[int32(v.ID)] = v
	}
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
	// 总分
	var sumScore int32 = 0
	// 写入作答记录
	records := make([]model.Record, 0, len(studentAnswer))
	for _, v := range studentAnswer {
		var lastScore int32 = 0
		if v.Answer == answerMap[v.QuesitonId].Answer {
			lastScore = questionScoreMap[v.QuesitonId]
		}
		sumScore += lastScore
		records = append(records, model.Record{
			StudentID:  int(userId),
			ExamID:     int(ExamId),
			PaperID:    paperInfo.ID,
			QuestionID: int(v.QuesitonId),
			Desc:       v.Answer,
			Score:      int(lastScore),
		})
	}
	if err := dal.CreateRecord(tx, &records); err != nil {
		tx.Rollback()
		return err
	}
	// 写入考试结果
	if err := dal.CreateExamResult(tx, &model.ExamResult{
		ExamID:    int(ExamId),
		StudentID: int(userId),
		Score:     int(sumScore),
	}); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

type StudentQeustionAnswer struct {
	QuesitonId int32  `json:"question_id"`
	Answer     string `json:"answer"`
}
