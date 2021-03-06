package dal

import (
	"gorm.io/gorm"
)

type Dal struct {
	db                     *gorm.DB
	teacherTableName       string
	studentTableName       string
	klassTableName         string
	teacherKlassTableName  string
	questionTableName      string
	paperTableName         string
	paperQuestionTableName string
	examTableName          string
	klassExamTableName     string
	examResultTableName    string
	recordTableName        string
}

var dal *Dal

func NewDal(gormDb *gorm.DB) {
	if dal == nil {
		dal = &Dal{
			db:                     gormDb,
			teacherTableName:       "teacher",
			studentTableName:       "student",
			klassTableName:         "klass",
			teacherKlassTableName:  "teacher_klass",
			questionTableName:      "question",
			paperTableName:         "paper",
			paperQuestionTableName: "paper_question",
			examTableName:          "exam",
			klassExamTableName:     "klass_exam",
			examResultTableName:    "exam_result",
			recordTableName:        "record",
		}
	}
}

func GetDb() *gorm.DB {
	return dal.db
}
