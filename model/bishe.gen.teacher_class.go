package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TeacherClassMgr struct {
	*_BaseMgr
}

// TeacherClassMgr open func
func TeacherClassMgr(db *gorm.DB) *_TeacherClassMgr {
	if db == nil {
		panic(fmt.Errorf("TeacherClassMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TeacherClassMgr{_BaseMgr: &_BaseMgr{DB: db.Table("teacher_class"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TeacherClassMgr) GetTableName() string {
	return "teacher_class"
}

// Get 获取
func (obj *_TeacherClassMgr) Get() (result TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TeacherClassMgr) Gets() (results []*TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增主键
func (obj *_TeacherClassMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTeacherID teacher_id获取 老师id
func (obj *_TeacherClassMgr) WithTeacherID(teacherID int) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// WithKlassID klass_id获取 班级id
func (obj *_TeacherClassMgr) WithKlassID(klassID int) Option {
	return optionFunc(func(o *options) { o.query["klass_id"] = klassID })
}

// GetByOption 功能选项模式获取
func (obj *_TeacherClassMgr) GetByOption(opts ...Option) (result TeacherClass, err error) {
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
func (obj *_TeacherClassMgr) GetByOptions(opts ...Option) (results []*TeacherClass, err error) {
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
func (obj *_TeacherClassMgr) GetFromID(id int) (result TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 自增主键
func (obj *_TeacherClassMgr) GetBatchFromID(ids []int) (results []*TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 老师id
func (obj *_TeacherClassMgr) GetFromTeacherID(teacherID int) (results []*TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`teacher_id` = ?", teacherID).Find(&results).Error

	return
}

// GetBatchFromTeacherID 批量查找 老师id
func (obj *_TeacherClassMgr) GetBatchFromTeacherID(teacherIDs []int) (results []*TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

// GetFromKlassID 通过klass_id获取内容 班级id
func (obj *_TeacherClassMgr) GetFromKlassID(klassID int) (results []*TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`klass_id` = ?", klassID).Find(&results).Error

	return
}

// GetBatchFromKlassID 批量查找 班级id
func (obj *_TeacherClassMgr) GetBatchFromKlassID(klassIDs []int) (results []*TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`klass_id` IN (?)", klassIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TeacherClassMgr) FetchByPrimaryKey(id int) (result TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUniTeacherIDKlassID primay or index 获取唯一内容
func (obj *_TeacherClassMgr) FetchUniqueIndexByUniTeacherIDKlassID(teacherID int, klassID int) (result TeacherClass, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`teacher_id` = ? AND `klass_id` = ?", teacherID, klassID).Find(&result).Error

	return
}
