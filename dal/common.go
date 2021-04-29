package dal

import (
	"fmt"

	"gorm.io/gorm"
)

type Dal struct {
	db *gorm.DB
}

var dal *Dal

func NewDal(gormDb *gorm.DB) {
	if dal == nil {
		dal = &Dal{
			db: gormDb,
		}
	}
}

func GetDb() (*gorm.DB, error) {
	if dal == nil {
		return nil, fmt.Errorf("GetDb() error : db is not exist")
	}
	if dal.db == nil {
		return nil, fmt.Errorf("GetDb() error : db is not exist")
	}
	return dal.db, nil
}
