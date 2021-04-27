package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RecordMgr struct {
	*_BaseMgr
}

// RecordMgr open func
func RecordMgr(db *gorm.DB) *_RecordMgr {
	if db == nil {
		panic(fmt.Errorf("RecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RecordMgr) GetTableName() string {
	return "record"
}

// Get 获取
func (obj *_RecordMgr) Get() (result Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RecordMgr) Gets() (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_RecordMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStudentID student_id获取 学生id
func (obj *_RecordMgr) WithStudentID(studentID int) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithExamID exam_id获取 考试id
func (obj *_RecordMgr) WithExamID(examID int) Option {
	return optionFunc(func(o *options) { o.query["exam_id"] = examID })
}

// WithPaperID paper_id获取 考卷id
func (obj *_RecordMgr) WithPaperID(paperID int) Option {
	return optionFunc(func(o *options) { o.query["paper_id"] = paperID })
}

// WithQuestionID question_id获取 题目id
func (obj *_RecordMgr) WithQuestionID(questionID int) Option {
	return optionFunc(func(o *options) { o.query["question_id"] = questionID })
}

// WithScore score获取 得分
func (obj *_RecordMgr) WithScore(score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// WithDesc desc获取 作答内容
func (obj *_RecordMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// GetByOption 功能选项模式获取
func (obj *_RecordMgr) GetByOption(opts ...Option) (result Record, err error) {
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
func (obj *_RecordMgr) GetByOptions(opts ...Option) (results []*Record, err error) {
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
func (obj *_RecordMgr) GetFromID(id int) (result Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_RecordMgr) GetBatchFromID(ids []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromStudentID 通过student_id获取内容 学生id
func (obj *_RecordMgr) GetFromStudentID(studentID int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`student_id` = ?", studentID).Find(&results).Error

	return
}

// GetBatchFromStudentID 批量查找 学生id
func (obj *_RecordMgr) GetBatchFromStudentID(studentIDs []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`student_id` IN (?)", studentIDs).Find(&results).Error

	return
}

// GetFromExamID 通过exam_id获取内容 考试id
func (obj *_RecordMgr) GetFromExamID(examID int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`exam_id` = ?", examID).Find(&results).Error

	return
}

// GetBatchFromExamID 批量查找 考试id
func (obj *_RecordMgr) GetBatchFromExamID(examIDs []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`exam_id` IN (?)", examIDs).Find(&results).Error

	return
}

// GetFromPaperID 通过paper_id获取内容 考卷id
func (obj *_RecordMgr) GetFromPaperID(paperID int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`paper_id` = ?", paperID).Find(&results).Error

	return
}

// GetBatchFromPaperID 批量查找 考卷id
func (obj *_RecordMgr) GetBatchFromPaperID(paperIDs []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`paper_id` IN (?)", paperIDs).Find(&results).Error

	return
}

// GetFromQuestionID 通过question_id获取内容 题目id
func (obj *_RecordMgr) GetFromQuestionID(questionID int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`question_id` = ?", questionID).Find(&results).Error

	return
}

// GetBatchFromQuestionID 批量查找 题目id
func (obj *_RecordMgr) GetBatchFromQuestionID(questionIDs []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`question_id` IN (?)", questionIDs).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 得分
func (obj *_RecordMgr) GetFromScore(score int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`score` = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量查找 得分
func (obj *_RecordMgr) GetBatchFromScore(scores []int) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`score` IN (?)", scores).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 作答内容
func (obj *_RecordMgr) GetFromDesc(desc string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 作答内容
func (obj *_RecordMgr) GetBatchFromDesc(descs []string) (results []*Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_RecordMgr) FetchByPrimaryKey(id int) (result Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUniKey primay or index 获取唯一内容
func (obj *_RecordMgr) FetchUniqueIndexByUniKey(studentID int, examID int, paperID int, questionID int) (result Record, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`student_id` = ? AND `exam_id` = ? AND `paper_id` = ? AND `question_id` = ?", studentID, examID, paperID, questionID).Find(&result).Error

	return
}
