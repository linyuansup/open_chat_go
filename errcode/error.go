package errcode

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Error struct {
	Code     int      `json:"code"`
	Message  string   `json:"message"`
	Detail   []string `json:"detail"`
	HttpCode int      `json:"httpCode"`
}

/**
400: 客户端可以处理的错误
501: 服务器内部错误
502: 鉴权失败强制拦截
*/

var codeList []int

func NewError(code int, message string, httpCode int) *Error {
	for _, e := range codeList {
		if e == code {
			panic("重复 Error code: " + fmt.Sprint(code))
		}
	}
	codeList = append(codeList, code)
	return &Error{Code: code, Message: message, HttpCode: httpCode}
}

func (e *Error) WithDetail(detail ...string) *Error {
	var newDetail []string
	for _, v := range e.Detail {
		newDetail = append(newDetail, strings.Clone(v))
	}
	newDetail = append(newDetail, detail...)
	return &Error{
		Code:     e.Code,
		Message:  strings.Clone(e.Message),
		Detail:   newDetail,
		HttpCode: e.HttpCode,
	}
}

func (e *Error) Error() string {
	marshal, _ := json.Marshal(e)
	return string(marshal)
}

func (e *Error) String() string {
	return "{\"code\":" + fmt.Sprint(e.Code) + "}"
}
