package model

import (
	"time"
)

// Exam [...]
type Exam struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null"`                          // 自增主键
	Name      string    `gorm:"column:name;type:varchar(255);not null"`                              // 考试名称
	BeginTime time.Time `gorm:"column:begin_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 考试结束时间
	EndTime   time.Time `gorm:"column:end_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`   // 考试结束时间
	PaperID   int       `gorm:"column:paper_id;type:int(11);not null"`                               // 考卷id
}

// ExamColumns get sql column name.获取数据库列名
var ExamColumns = struct {
	ID        string
	Name      string
	BeginTime string
	EndTime   string
	PaperID   string
}{
	ID:        "id",
	Name:      "name",
	BeginTime: "begin_time",
	EndTime:   "end_time",
	PaperID:   "paper_id",
}

// ExamResult [...]
type ExamResult struct {
	ID        int `gorm:"primaryKey;column:id;type:int(11);not null"`                                 // 自增主键
	ExamID    int `gorm:"uniqueIndex:uni_exam_id_student_id;column:exam_id;type:int(11);not null"`    // 考试id
	StudentID int `gorm:"uniqueIndex:uni_exam_id_student_id;column:student_id;type:int(11);not null"` // 学生id
	Score     int `gorm:"column:score;type:int(11);not null"`                                         // 得分
}

// ExamResultColumns get sql column name.获取数据库列名
var ExamResultColumns = struct {
	ID        string
	ExamID    string
	StudentID string
	Score     string
}{
	ID:        "id",
	ExamID:    "exam_id",
	StudentID: "student_id",
	Score:     "score",
}

// Klass [...]
type Klass struct {
	ID   int    `gorm:"primaryKey;column:id;type:int(11);not null"` // 自增主键
	Name string `gorm:"column:name;type:varchar(255);not null"`     // 名称
}

// KlassColumns get sql column name.获取数据库列名
var KlassColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

// KlassExam [...]
type KlassExam struct {
	ID      int `gorm:"primaryKey;column:id;type:int(11);not null"`                             // 自增主键
	KlassID int `gorm:"uniqueIndex:uni_klass_id_exam_id;column:klass_id;type:int(11);not null"` // 班级id
	ExamID  int `gorm:"uniqueIndex:uni_klass_id_exam_id;column:exam_id;type:int(11);not null"`  // 考试id
}

// KlassExamColumns get sql column name.获取数据库列名
var KlassExamColumns = struct {
	ID      string
	KlassID string
	ExamID  string
}{
	ID:      "id",
	KlassID: "klass_id",
	ExamID:  "exam_id",
}

// Paper [...]
type Paper struct {
	ID         int    `gorm:"primaryKey;column:id;type:int(11);not null"` // 自增主键
	Name       string `gorm:"column:name;type:varchar(255);not null"`     // 考卷名称
	ScoreLimit int    `gorm:"column:score_limit;type:int(11);not null"`   // 分数上限
}

// PaperColumns get sql column name.获取数据库列名
var PaperColumns = struct {
	ID         string
	Name       string
	ScoreLimit string
}{
	ID:         "id",
	Name:       "name",
	ScoreLimit: "score_limit",
}

// PaperQuestion [...]
type PaperQuestion struct {
	ID            int `gorm:"primaryKey;column:id;type:int(11);not null"`                                    // 自增主键
	PaperID       int `gorm:"uniqueIndex:uni_paper_id_question_id;column:paper_id;type:int(11);not null"`    // 考卷id
	QuestionID    int `gorm:"uniqueIndex:uni_paper_id_question_id;column:question_id;type:int(11);not null"` // 题目id
	QuestionScore int `gorm:"column:question_score;type:int(11);not null"`                                   // 题目得分
}

// PaperQuestionColumns get sql column name.获取数据库列名
var PaperQuestionColumns = struct {
	ID            string
	PaperID       string
	QuestionID    string
	QuestionScore string
}{
	ID:            "id",
	PaperID:       "paper_id",
	QuestionID:    "question_id",
	QuestionScore: "question_score",
}

// Question [...]
type Question struct {
	ID          int    `gorm:"primaryKey;column:id;type:int(11);not null"` // 自增主键
	Desc        string `gorm:"column:desc;type:varchar(255);not null"`     // 题目描述
	Answer      string `gorm:"column:answer;type:varchar(255);not null"`   // 题目答案
	Type        int    `gorm:"column:type;type:int(11);not null"`          // 题目类型：0 未知，1 单选，2多选
	OptionDescA string `gorm:"column:option_desc_A;type:varchar(255)"`     // 选项 A，选项内容为空，即没有该选项
	OptionDescB string `gorm:"column:option_desc_B;type:varchar(255)"`     // 选项 B
	OptionDescC string `gorm:"column:option_desc_C;type:varchar(255)"`     // 选项 C
	OptionDescD string `gorm:"column:option_desc_D;type:varchar(255)"`     // 选项 D
}

// QuestionColumns get sql column name.获取数据库列名
var QuestionColumns = struct {
	ID          string
	Desc        string
	Answer      string
	Type        string
	OptionDescA string
	OptionDescB string
	OptionDescC string
	OptionDescD string
}{
	ID:          "id",
	Desc:        "desc",
	Answer:      "answer",
	Type:        "type",
	OptionDescA: "option_desc_A",
	OptionDescB: "option_desc_B",
	OptionDescC: "option_desc_C",
	OptionDescD: "option_desc_D",
}

// Record [...]
type Record struct {
	ID         int    `gorm:"primaryKey;column:id;type:int(11);not null"`                   // 自增主键
	StudentID  int    `gorm:"uniqueIndex:uni_key;column:student_id;type:int(11);not null"`  // 学生id
	ExamID     int    `gorm:"uniqueIndex:uni_key;column:exam_id;type:int(11);not null"`     // 考试id
	PaperID    int    `gorm:"uniqueIndex:uni_key;column:paper_id;type:int(11);not null"`    // 考卷id
	QuestionID int    `gorm:"uniqueIndex:uni_key;column:question_id;type:int(11);not null"` // 题目id
	Score      int    `gorm:"column:score;type:int(11);not null"`                           // 得分
	Desc       string `gorm:"column:desc;type:varchar(255)"`                                // 作答内容
}

// RecordColumns get sql column name.获取数据库列名
var RecordColumns = struct {
	ID         string
	StudentID  string
	ExamID     string
	PaperID    string
	QuestionID string
	Score      string
	Desc       string
}{
	ID:         "id",
	StudentID:  "student_id",
	ExamID:     "exam_id",
	PaperID:    "paper_id",
	QuestionID: "question_id",
	Score:      "score",
	Desc:       "desc",
}

// Student [...]
type Student struct {
	ID       int    `gorm:"primaryKey;column:id;type:int(11);not null"`      // 自增主键
	Name     string `gorm:"column:name;type:varchar(255);not null"`          // 名称
	Number   string `gorm:"unique;column:number;type:varchar(255);not null"` // 学号
	Password string `gorm:"column:password;type:varchar(255);not null"`      // 密码
	KlassID  int    `gorm:"column:klass_id;type:int(11);not null"`           // 所属班级id
}

// StudentColumns get sql column name.获取数据库列名
var StudentColumns = struct {
	ID       string
	Name     string
	Number   string
	Password string
	KlassID  string
}{
	ID:       "id",
	Name:     "name",
	Number:   "number",
	Password: "password",
	KlassID:  "klass_id",
}

// Teacher [...]
type Teacher struct {
	ID       int    `gorm:"primaryKey;column:id;type:int(11);not null"`      // 自增主键
	Name     string `gorm:"column:name;type:varchar(255);not null"`          // 名称
	Number   string `gorm:"unique;column:number;type:varchar(255);not null"` // 工号
	Password string `gorm:"column:password;type:varchar(255);not null"`      // 密码
}

// TeacherColumns get sql column name.获取数据库列名
var TeacherColumns = struct {
	ID       string
	Name     string
	Number   string
	Password string
}{
	ID:       "id",
	Name:     "name",
	Number:   "number",
	Password: "password",
}

// TeacherKlass [...]
type TeacherKlass struct {
	ID        int `gorm:"primaryKey;column:id;type:int(11);not null"`                                  // 自增主键
	TeacherID int `gorm:"uniqueIndex:uni_teacher_id_klass_id;column:teacher_id;type:int(11);not null"` // 老师id
	KlassID   int `gorm:"uniqueIndex:uni_teacher_id_klass_id;column:klass_id;type:int(11);not null"`   // 班级id
}

// TeacherKlassColumns get sql column name.获取数据库列名
var TeacherKlassColumns = struct {
	ID        string
	TeacherID string
	KlassID   string
}{
	ID:        "id",
	TeacherID: "teacher_id",
	KlassID:   "klass_id",
}
