package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _StudentMgr struct {
	*_BaseMgr
}

// StudentMgr open func
func StudentMgr(db *gorm.DB) *_StudentMgr {
	if db == nil {
		panic(fmt.Errorf("StudentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StudentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("student"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StudentMgr) GetTableName() string {
	return "student"
}

// Get 获取
func (obj *_StudentMgr) Get() (result Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StudentMgr) Gets() (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_StudentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_StudentMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithNumber number获取 学号
func (obj *_StudentMgr) WithNumber(number string) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithPassword password获取 密码
func (obj *_StudentMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithKlassID klass_id获取 所属班级id
func (obj *_StudentMgr) WithKlassID(klassID int) Option {
	return optionFunc(func(o *options) { o.query["klass_id"] = klassID })
}

// GetByOption 功能选项模式获取
func (obj *_StudentMgr) GetByOption(opts ...Option) (result Student, err error) {
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
func (obj *_StudentMgr) GetByOptions(opts ...Option) (results []*Student, err error) {
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
func (obj *_StudentMgr) GetFromID(id int) (result Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_StudentMgr) GetBatchFromID(ids []int) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_StudentMgr) GetFromName(name string) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_StudentMgr) GetBatchFromName(names []string) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容 学号
func (obj *_StudentMgr) GetFromNumber(number string) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`number` = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量查找 学号
func (obj *_StudentMgr) GetBatchFromNumber(numbers []string) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`number` IN (?)", numbers).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_StudentMgr) GetFromPassword(password string) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_StudentMgr) GetBatchFromPassword(passwords []string) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromKlassID 通过klass_id获取内容 所属班级id
func (obj *_StudentMgr) GetFromKlassID(klassID int) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`klass_id` = ?", klassID).Find(&results).Error

	return
}

// GetBatchFromKlassID 批量查找 所属班级id
func (obj *_StudentMgr) GetBatchFromKlassID(klassIDs []int) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`klass_id` IN (?)", klassIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_StudentMgr) FetchByPrimaryKey(id int) (result Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
