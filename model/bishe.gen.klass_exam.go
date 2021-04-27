package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _KlassExamMgr struct {
	*_BaseMgr
}

// KlassExamMgr open func
func KlassExamMgr(db *gorm.DB) *_KlassExamMgr {
	if db == nil {
		panic(fmt.Errorf("KlassExamMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_KlassExamMgr{_BaseMgr: &_BaseMgr{DB: db.Table("klass_exam"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_KlassExamMgr) GetTableName() string {
	return "klass_exam"
}

// Get 获取
func (obj *_KlassExamMgr) Get() (result KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_KlassExamMgr) Gets() (results []*KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_KlassExamMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithKlassID klass_id获取 班级id
func (obj *_KlassExamMgr) WithKlassID(klassID int) Option {
	return optionFunc(func(o *options) { o.query["klass_id"] = klassID })
}

// WithExamID exam_id获取 考试id
func (obj *_KlassExamMgr) WithExamID(examID int) Option {
	return optionFunc(func(o *options) { o.query["exam_id"] = examID })
}

// GetByOption 功能选项模式获取
func (obj *_KlassExamMgr) GetByOption(opts ...Option) (result KlassExam, err error) {
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
func (obj *_KlassExamMgr) GetByOptions(opts ...Option) (results []*KlassExam, err error) {
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
func (obj *_KlassExamMgr) GetFromID(id int) (result KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_KlassExamMgr) GetBatchFromID(ids []int) (results []*KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromKlassID 通过klass_id获取内容 班级id
func (obj *_KlassExamMgr) GetFromKlassID(klassID int) (results []*KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`klass_id` = ?", klassID).Find(&results).Error

	return
}

// GetBatchFromKlassID 批量查找 班级id
func (obj *_KlassExamMgr) GetBatchFromKlassID(klassIDs []int) (results []*KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`klass_id` IN (?)", klassIDs).Find(&results).Error

	return
}

// GetFromExamID 通过exam_id获取内容 考试id
func (obj *_KlassExamMgr) GetFromExamID(examID int) (results []*KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`exam_id` = ?", examID).Find(&results).Error

	return
}

// GetBatchFromExamID 批量查找 考试id
func (obj *_KlassExamMgr) GetBatchFromExamID(examIDs []int) (results []*KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`exam_id` IN (?)", examIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_KlassExamMgr) FetchByPrimaryKey(id int) (result KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUniKlassIDExamID primay or index 获取唯一内容
func (obj *_KlassExamMgr) FetchUniqueIndexByUniKlassIDExamID(klassID int, examID int) (result KlassExam, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`klass_id` = ? AND `exam_id` = ?", klassID, examID).Find(&result).Error

	return
}
