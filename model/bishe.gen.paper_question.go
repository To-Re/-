package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PaperQuestionMgr struct {
	*_BaseMgr
}

// PaperQuestionMgr open func
func PaperQuestionMgr(db *gorm.DB) *_PaperQuestionMgr {
	if db == nil {
		panic(fmt.Errorf("PaperQuestionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PaperQuestionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("paper_question"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PaperQuestionMgr) GetTableName() string {
	return "paper_question"
}

// Get 获取
func (obj *_PaperQuestionMgr) Get() (result PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PaperQuestionMgr) Gets() (results []*PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_PaperQuestionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPaperID paper_id获取 考卷id
func (obj *_PaperQuestionMgr) WithPaperID(paperID int) Option {
	return optionFunc(func(o *options) { o.query["paper_id"] = paperID })
}

// WithQuestionID question_id获取 题目id
func (obj *_PaperQuestionMgr) WithQuestionID(questionID int) Option {
	return optionFunc(func(o *options) { o.query["question_id"] = questionID })
}

// GetByOption 功能选项模式获取
func (obj *_PaperQuestionMgr) GetByOption(opts ...Option) (result PaperQuestion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PaperQuestionMgr) GetByOptions(opts ...Option) (results []*PaperQuestion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 自增主键
func (obj *_PaperQuestionMgr) GetFromID(id int) (result PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_PaperQuestionMgr) GetBatchFromID(ids []int) (results []*PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPaperID 通过paper_id获取内容 考卷id
func (obj *_PaperQuestionMgr) GetFromPaperID(paperID int) (results []*PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`paper_id` = ?", paperID).Find(&results).Error

	return
}

// GetBatchFromPaperID 批量查找 考卷id
func (obj *_PaperQuestionMgr) GetBatchFromPaperID(paperIDs []int) (results []*PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`paper_id` IN (?)", paperIDs).Find(&results).Error

	return
}

// GetFromQuestionID 通过question_id获取内容 题目id
func (obj *_PaperQuestionMgr) GetFromQuestionID(questionID int) (results []*PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`question_id` = ?", questionID).Find(&results).Error

	return
}

// GetBatchFromQuestionID 批量查找 题目id
func (obj *_PaperQuestionMgr) GetBatchFromQuestionID(questionIDs []int) (results []*PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`question_id` IN (?)", questionIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PaperQuestionMgr) FetchByPrimaryKey(id int) (result PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUniPaperIDQuestionID primay or index 获取唯一内容
func (obj *_PaperQuestionMgr) FetchUniqueIndexByUniPaperIDQuestionID(paperID int, questionID int) (result PaperQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`paper_id` = ? AND `question_id` = ?", paperID, questionID).Find(&result).Error

	return
}
