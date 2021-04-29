package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TeacherKlassMgr struct {
	*_BaseMgr
}

// TeacherKlassMgr open func
func TeacherKlassMgr(db *gorm.DB) *_TeacherKlassMgr {
	if db == nil {
		panic(fmt.Errorf("TeacherKlassMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TeacherKlassMgr{_BaseMgr: &_BaseMgr{DB: db.Table("teacher_klass"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TeacherKlassMgr) GetTableName() string {
	return "teacher_klass"
}

// Get 获取
func (obj *_TeacherKlassMgr) Get() (result TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TeacherKlassMgr) Gets() (results []*TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_TeacherKlassMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTeacherID teacher_id获取 老师id
func (obj *_TeacherKlassMgr) WithTeacherID(teacherID int) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// WithKlassID klass_id获取 班级id
func (obj *_TeacherKlassMgr) WithKlassID(klassID int) Option {
	return optionFunc(func(o *options) { o.query["klass_id"] = klassID })
}

// GetByOption 功能选项模式获取
func (obj *_TeacherKlassMgr) GetByOption(opts ...Option) (result TeacherKlass, err error) {
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
func (obj *_TeacherKlassMgr) GetByOptions(opts ...Option) (results []*TeacherKlass, err error) {
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
func (obj *_TeacherKlassMgr) GetFromID(id int) (result TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_TeacherKlassMgr) GetBatchFromID(ids []int) (results []*TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 老师id
func (obj *_TeacherKlassMgr) GetFromTeacherID(teacherID int) (results []*TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`teacher_id` = ?", teacherID).Find(&results).Error

	return
}

// GetBatchFromTeacherID 批量查找 老师id
func (obj *_TeacherKlassMgr) GetBatchFromTeacherID(teacherIDs []int) (results []*TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

// GetFromKlassID 通过klass_id获取内容 班级id
func (obj *_TeacherKlassMgr) GetFromKlassID(klassID int) (results []*TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`klass_id` = ?", klassID).Find(&results).Error

	return
}

// GetBatchFromKlassID 批量查找 班级id
func (obj *_TeacherKlassMgr) GetBatchFromKlassID(klassIDs []int) (results []*TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`klass_id` IN (?)", klassIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TeacherKlassMgr) FetchByPrimaryKey(id int) (result TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUniTeacherIDKlassID primay or index 获取唯一内容
func (obj *_TeacherKlassMgr) FetchUniqueIndexByUniTeacherIDKlassID(teacherID int, klassID int) (result TeacherKlass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`teacher_id` = ? AND `klass_id` = ?", teacherID, klassID).Find(&result).Error

	return
}
