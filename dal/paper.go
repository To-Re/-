package dal

import (
	"bishe/backend/model"

	"gorm.io/gorm"
)

func GetPaperList() ([]*model.Paper, error) {
	db := dal.db
	list := []*model.Paper{}
	if err := db.Table(dal.paperTableName).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func CreatePaper(paper *model.Paper) error {
	db := dal.db
	if err := db.Table(dal.paperTableName).Create(paper).Error; err != nil {
		return err
	}
	return nil
}

func GetPaperById(paper_id int32) (*model.Paper, error) {
	db := dal.db
	paper := &model.Paper{}
	if err := db.Table(dal.paperTableName).Where("id = ?", paper_id).First(&paper).Error; err != nil {
		return nil, err
	}
	return paper, nil
}

func UpdatePaperName(paper *model.Paper) error {
	db := dal.db
	if err := db.Table(dal.paperTableName).
		Where("id = ?", paper.ID).
		Select("name").
		Updates(paper).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePaperScore(tx *gorm.DB, paper *model.Paper) error {
	if err := tx.Table(dal.paperTableName).
		Where("id = ?", paper.ID).
		Select("score_limit").
		Updates(paper).Error; err != nil {
		return err
	}
	return nil
}
