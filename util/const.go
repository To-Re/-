package util

const UserTypeTeacher = 1
const UserTypeStudent = 2

const QuestionTypeDanXuan = 1
const QuestionTypeDuoXuan = 2

var QuestionTypeMap = map[int32]string{
	QuestionTypeDanXuan: "单选题",
	QuestionTypeDuoXuan: "多选题",
}

const TimeTemplate = "2006-01-02 15:04:05"
