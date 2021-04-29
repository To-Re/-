package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TeacherMgr struct {
	*_BaseMgr
}

// TeacherMgr open func
func TeacherMgr(db *gorm.DB) *_TeacherMgr {
	if db == nil {
		panic(fmt.Errorf("TeacherMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TeacherMgr{_BaseMgr: &_BaseMgr{DB: db.Table("teacher"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TeacherMgr) GetTableName() string {
	return "teacher"
}

// Get 获取
func (obj *_TeacherMgr) Get() (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TeacherMgr) Gets() (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_TeacherMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_TeacherMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithNumber number获取 工号
func (obj *_TeacherMgr) WithNumber(number string) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithPassword password获取 密码
func (obj *_TeacherMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// GetByOption 功能选项模式获取
func (obj *_TeacherMgr) GetByOption(opts ...Option) (result Teacher, err error) {
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
func (obj *_TeacherMgr) GetByOptions(opts ...Option) (results []*Teacher, err error) {
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
func (obj *_TeacherMgr) GetFromID(id int) (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_TeacherMgr) GetBatchFromID(ids []int) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_TeacherMgr) GetFromName(name string) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_TeacherMgr) GetBatchFromName(names []string) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容 工号
func (obj *_TeacherMgr) GetFromNumber(number string) (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`number` = ?", number).Find(&result).Error

	return
}

// GetBatchFromNumber 批量查找 工号
func (obj *_TeacherMgr) GetBatchFromNumber(numbers []string) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`number` IN (?)", numbers).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_TeacherMgr) GetFromPassword(password string) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_TeacherMgr) GetBatchFromPassword(passwords []string) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TeacherMgr) FetchByPrimaryKey(id int) (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByNumber primay or index 获取唯一内容
func (obj *_TeacherMgr) FetchUniqueByNumber(number string) (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`number` = ?", number).Find(&result).Error

	return
}
