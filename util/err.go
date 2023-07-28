package util

import "errors"

// ErrCode 错误代码
type ErrCode int

const (
	CommonError     ErrCode = -1 // 公共错误码，表示接口错误
	PageNotFound    ErrCode = 404
	SystemAbnormal  ErrCode = 500
	UnknownAbnormal ErrCode = 1000
	DbAbnormal      ErrCode = 1001
	ReqBindError    ErrCode = 1010
	ReqParamsError  ErrCode = 1011
)

func (e ErrCode) ToInt() int {
	return int(e)
}

func (e ErrCode) NewError() error {
	return errors.New(e.GetText())
}

func (e ErrCode) GetText() string {
	textMap := map[ErrCode]string{
		PageNotFound:    "请求资源不存在",
		SystemAbnormal:  "系统异常，请稍候重试",
		UnknownAbnormal: "未知异常，请稍候重试",
		DbAbnormal:      "数据库异常，请稍候重试",
		ReqBindError:    "请求数据解析错误",
		ReqParamsError:  "请求数据错误",
	}

	text, ok := textMap[e]
	if !ok {
		return textMap[UnknownAbnormal]
	}

	return text
}
