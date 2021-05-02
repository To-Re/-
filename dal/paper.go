package dal

import "bishe/backend/model"

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
