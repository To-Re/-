package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _QuestionMgr struct {
	*_BaseMgr
}

// QuestionMgr open func
func QuestionMgr(db *gorm.DB) *_QuestionMgr {
	if db == nil {
		panic(fmt.Errorf("QuestionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_QuestionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("question"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_QuestionMgr) GetTableName() string {
	return "question"
}

// Get 获取
func (obj *_QuestionMgr) Get() (result Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_QuestionMgr) Gets() (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_QuestionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDesc desc获取 题目描述
func (obj *_QuestionMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithAnswer answer获取 题目答案
func (obj *_QuestionMgr) WithAnswer(answer string) Option {
	return optionFunc(func(o *options) { o.query["answer"] = answer })
}

// WithType type获取 题目类型：0 未知，1 单选，2多选
func (obj *_QuestionMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithOptionDescA option_desc_A获取 选项 A，选项内容为空，即没有该选项
func (obj *_QuestionMgr) WithOptionDescA(optionDescA string) Option {
	return optionFunc(func(o *options) { o.query["option_desc_A"] = optionDescA })
}

// WithOptionDescB option_desc_B获取 选项 B
func (obj *_QuestionMgr) WithOptionDescB(optionDescB string) Option {
	return optionFunc(func(o *options) { o.query["option_desc_B"] = optionDescB })
}

// WithOptionDescC option_desc_C获取 选项 C
func (obj *_QuestionMgr) WithOptionDescC(optionDescC string) Option {
	return optionFunc(func(o *options) { o.query["option_desc_C"] = optionDescC })
}

// WithOptionDescD option_desc_D获取 选项 D
func (obj *_QuestionMgr) WithOptionDescD(optionDescD string) Option {
	return optionFunc(func(o *options) { o.query["option_desc_D"] = optionDescD })
}

// GetByOption 功能选项模式获取
func (obj *_QuestionMgr) GetByOption(opts ...Option) (result Question, err error) {
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
func (obj *_QuestionMgr) GetByOptions(opts ...Option) (results []*Question, err error) {
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
func (obj *_QuestionMgr) GetFromID(id int) (result Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_QuestionMgr) GetBatchFromID(ids []int) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 题目描述
func (obj *_QuestionMgr) GetFromDesc(desc string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 题目描述
func (obj *_QuestionMgr) GetBatchFromDesc(descs []string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

// GetFromAnswer 通过answer获取内容 题目答案
func (obj *_QuestionMgr) GetFromAnswer(answer string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`answer` = ?", answer).Find(&results).Error

	return
}

// GetBatchFromAnswer 批量查找 题目答案
func (obj *_QuestionMgr) GetBatchFromAnswer(answers []string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`answer` IN (?)", answers).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 题目类型：0 未知，1 单选，2多选
func (obj *_QuestionMgr) GetFromType(_type int) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 题目类型：0 未知，1 单选，2多选
func (obj *_QuestionMgr) GetBatchFromType(_types []int) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromOptionDescA 通过option_desc_A获取内容 选项 A，选项内容为空，即没有该选项
func (obj *_QuestionMgr) GetFromOptionDescA(optionDescA string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`option_desc_A` = ?", optionDescA).Find(&results).Error

	return
}

// GetBatchFromOptionDescA 批量查找 选项 A，选项内容为空，即没有该选项
func (obj *_QuestionMgr) GetBatchFromOptionDescA(optionDescAs []string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`option_desc_A` IN (?)", optionDescAs).Find(&results).Error

	return
}

// GetFromOptionDescB 通过option_desc_B获取内容 选项 B
func (obj *_QuestionMgr) GetFromOptionDescB(optionDescB string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`option_desc_B` = ?", optionDescB).Find(&results).Error

	return
}

// GetBatchFromOptionDescB 批量查找 选项 B
func (obj *_QuestionMgr) GetBatchFromOptionDescB(optionDescBs []string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`option_desc_B` IN (?)", optionDescBs).Find(&results).Error

	return
}

// GetFromOptionDescC 通过option_desc_C获取内容 选项 C
func (obj *_QuestionMgr) GetFromOptionDescC(optionDescC string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`option_desc_C` = ?", optionDescC).Find(&results).Error

	return
}

// GetBatchFromOptionDescC 批量查找 选项 C
func (obj *_QuestionMgr) GetBatchFromOptionDescC(optionDescCs []string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`option_desc_C` IN (?)", optionDescCs).Find(&results).Error

	return
}

// GetFromOptionDescD 通过option_desc_D获取内容 选项 D
func (obj *_QuestionMgr) GetFromOptionDescD(optionDescD string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`option_desc_D` = ?", optionDescD).Find(&results).Error

	return
}

// GetBatchFromOptionDescD 批量查找 选项 D
func (obj *_QuestionMgr) GetBatchFromOptionDescD(optionDescDs []string) (results []*Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`option_desc_D` IN (?)", optionDescDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_QuestionMgr) FetchByPrimaryKey(id int) (result Question, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
