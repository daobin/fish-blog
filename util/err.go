package util

import "errors"

// ErrCode 错误代码
type ErrCode int

const (
	ErrCodeSystem  ErrCode = 500
	ErrCodeUnknown ErrCode = 1000
	ErrCodeDb      ErrCode = 1001
	ErrCodeReqBind ErrCode = 1010
)

func (e ErrCode) ToInt() int {
	return int(e)
}

func (e ErrCode) NewError() error {
	return errors.New(e.GetText())
}

func (e ErrCode) GetText() string {
	textMap := map[ErrCode]string{
		ErrCodeSystem:  "系统异常，请稍候重试",
		ErrCodeUnknown: "未知异常，请稍候重试",
		ErrCodeDb:      "数据库异常，请稍候重试",
		ErrCodeReqBind: "请求数据解析错误",
	}

	text, ok := textMap[e]
	if !ok {
		return textMap[ErrCodeUnknown]
	}

	return text
}
