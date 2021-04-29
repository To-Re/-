package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ExamMgr struct {
	*_BaseMgr
}

// ExamMgr open func
func ExamMgr(db *gorm.DB) *_ExamMgr {
	if db == nil {
		panic(fmt.Errorf("ExamMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ExamMgr{_BaseMgr: &_BaseMgr{DB: db.Table("exam"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ExamMgr) GetTableName() string {
	return "exam"
}

// Get 获取
func (obj *_ExamMgr) Get() (result Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ExamMgr) Gets() (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_ExamMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 考试名称
func (obj *_ExamMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithBeginTime begin_time获取 考试开始时间
func (obj *_ExamMgr) WithBeginTime(beginTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["begin_time"] = beginTime })
}

// WithLength length获取 考试时长
func (obj *_ExamMgr) WithLength(length int) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithPaperID paper_id获取 考卷id
func (obj *_ExamMgr) WithPaperID(paperID int) Option {
	return optionFunc(func(o *options) { o.query["paper_id"] = paperID })
}

// GetByOption 功能选项模式获取
func (obj *_ExamMgr) GetByOption(opts ...Option) (result Exam, err error) {
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
func (obj *_ExamMgr) GetByOptions(opts ...Option) (results []*Exam, err error) {
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
func (obj *_ExamMgr) GetFromID(id int) (result Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_ExamMgr) GetBatchFromID(ids []int) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 考试名称
func (obj *_ExamMgr) GetFromName(name string) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 考试名称
func (obj *_ExamMgr) GetBatchFromName(names []string) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromBeginTime 通过begin_time获取内容 考试开始时间
func (obj *_ExamMgr) GetFromBeginTime(beginTime time.Time) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`begin_time` = ?", beginTime).Find(&results).Error

	return
}

// GetBatchFromBeginTime 批量查找 考试开始时间
func (obj *_ExamMgr) GetBatchFromBeginTime(beginTimes []time.Time) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`begin_time` IN (?)", beginTimes).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容 考试时长
func (obj *_ExamMgr) GetFromLength(length int) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找 考试时长
func (obj *_ExamMgr) GetBatchFromLength(lengths []int) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromPaperID 通过paper_id获取内容 考卷id
func (obj *_ExamMgr) GetFromPaperID(paperID int) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`paper_id` = ?", paperID).Find(&results).Error

	return
}

// GetBatchFromPaperID 批量查找 考卷id
func (obj *_ExamMgr) GetBatchFromPaperID(paperIDs []int) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`paper_id` IN (?)", paperIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ExamMgr) FetchByPrimaryKey(id int) (result Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
