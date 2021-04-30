package util

type ErrorInfo struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func BuildError(code int32, msg string) *ErrorInfo {
	return &ErrorInfo{
		Code: code,
		Msg:  msg,
	}
}

const OK = 0
const NOTLOGIN = -1
const PARAMERROR = -2
const NETWORKERROR = -3
const NOTAUTH = -4
const LOGINERROR = -100
const REGISTERERROR = -101

var ErrMap = map[int32]string{
	OK:            "成功",
	NOTLOGIN:      "未登录",
	PARAMERROR:    "参数错误",
	NETWORKERROR:  "网络错误",
	LOGINERROR:    "登录失败",
	NOTAUTH:       "权限不足",
	REGISTERERROR: "注册失败",
}
