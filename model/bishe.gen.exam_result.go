package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ExamResultMgr struct {
	*_BaseMgr
}

// ExamResultMgr open func
func ExamResultMgr(db *gorm.DB) *_ExamResultMgr {
	if db == nil {
		panic(fmt.Errorf("ExamResultMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ExamResultMgr{_BaseMgr: &_BaseMgr{DB: db.Table("exam_result"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ExamResultMgr) GetTableName() string {
	return "exam_result"
}

// Get 获取
func (obj *_ExamResultMgr) Get() (result ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ExamResultMgr) Gets() (results []*ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_ExamResultMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithExamID exam_id获取 考试id
func (obj *_ExamResultMgr) WithExamID(examID int) Option {
	return optionFunc(func(o *options) { o.query["exam_id"] = examID })
}

// WithStudentID student_id获取 学生id
func (obj *_ExamResultMgr) WithStudentID(studentID int) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithScore score获取 得分
func (obj *_ExamResultMgr) WithScore(score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// GetByOption 功能选项模式获取
func (obj *_ExamResultMgr) GetByOption(opts ...Option) (result ExamResult, err error) {
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
func (obj *_ExamResultMgr) GetByOptions(opts ...Option) (results []*ExamResult, err error) {
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
func (obj *_ExamResultMgr) GetFromID(id int) (result ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_ExamResultMgr) GetBatchFromID(ids []int) (results []*ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromExamID 通过exam_id获取内容 考试id
func (obj *_ExamResultMgr) GetFromExamID(examID int) (results []*ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`exam_id` = ?", examID).Find(&results).Error

	return
}

// GetBatchFromExamID 批量查找 考试id
func (obj *_ExamResultMgr) GetBatchFromExamID(examIDs []int) (results []*ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`exam_id` IN (?)", examIDs).Find(&results).Error

	return
}

// GetFromStudentID 通过student_id获取内容 学生id
func (obj *_ExamResultMgr) GetFromStudentID(studentID int) (results []*ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`student_id` = ?", studentID).Find(&results).Error

	return
}

// GetBatchFromStudentID 批量查找 学生id
func (obj *_ExamResultMgr) GetBatchFromStudentID(studentIDs []int) (results []*ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`student_id` IN (?)", studentIDs).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 得分
func (obj *_ExamResultMgr) GetFromScore(score int) (results []*ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`score` = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量查找 得分
func (obj *_ExamResultMgr) GetBatchFromScore(scores []int) (results []*ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`score` IN (?)", scores).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ExamResultMgr) FetchByPrimaryKey(id int) (result ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUniExamIDStudentID primay or index 获取唯一内容
func (obj *_ExamResultMgr) FetchUniqueIndexByUniExamIDStudentID(examID int, studentID int) (result ExamResult, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`exam_id` = ? AND `student_id` = ?", examID, studentID).Find(&result).Error

	return
}
