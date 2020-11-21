package errors

import (
	"encoding/json"
)

// Err 异常结构体
type Err struct {
	Code int
	Msg  string
}

// Error 实现异常接口
func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

// New 新建异常
func New(code int, msg string) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}
