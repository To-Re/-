package util

type error_info struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func BuildError(code int32, msg string) *error_info {
	return &error_info{
		Code: code,
		Msg:  msg,
	}
}

const NOTLOGIN = -1

var ErrMap = map[int32]string{
	NOTLOGIN: "未登录",
}
