package dal

import (
	"gorm.io/gorm"
)

type Dal struct {
	db               *gorm.DB
	teacherTableName string
}

var dal *Dal

func NewDal(gormDb *gorm.DB) {
	if dal == nil {
		dal = &Dal{
			db:               gormDb,
			teacherTableName: "teacher",
		}
	}
}
