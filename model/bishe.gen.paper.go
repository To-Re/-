package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PaperMgr struct {
	*_BaseMgr
}

// PaperMgr open func
func PaperMgr(db *gorm.DB) *_PaperMgr {
	if db == nil {
		panic(fmt.Errorf("PaperMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PaperMgr{_BaseMgr: &_BaseMgr{DB: db.Table("paper"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PaperMgr) GetTableName() string {
	return "paper"
}

// Get 获取
func (obj *_PaperMgr) Get() (result Paper, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PaperMgr) Gets() (results []*Paper, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_PaperMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 考卷名称
func (obj *_PaperMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithScoreLimit score_limit获取 分数上限
func (obj *_PaperMgr) WithScoreLimit(scoreLimit int) Option {
	return optionFunc(func(o *options) { o.query["score_limit"] = scoreLimit })
}

// GetByOption 功能选项模式获取
func (obj *_PaperMgr) GetByOption(opts ...Option) (result Paper, err error) {
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
func (obj *_PaperMgr) GetByOptions(opts ...Option) (results []*Paper, err error) {
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
func (obj *_PaperMgr) GetFromID(id int) (result Paper, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_PaperMgr) GetBatchFromID(ids []int) (results []*Paper, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 考卷名称
func (obj *_PaperMgr) GetFromName(name string) (results []*Paper, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 考卷名称
func (obj *_PaperMgr) GetBatchFromName(names []string) (results []*Paper, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromScoreLimit 通过score_limit获取内容 分数上限
func (obj *_PaperMgr) GetFromScoreLimit(scoreLimit int) (results []*Paper, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`score_limit` = ?", scoreLimit).Find(&results).Error

	return
}

// GetBatchFromScoreLimit 批量查找 分数上限
func (obj *_PaperMgr) GetBatchFromScoreLimit(scoreLimits []int) (results []*Paper, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`score_limit` IN (?)", scoreLimits).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PaperMgr) FetchByPrimaryKey(id int) (result Paper, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
