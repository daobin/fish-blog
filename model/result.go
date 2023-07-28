package model

import "net/http"

// Result HTTP接口响应结构体
type Result[T any] struct {
	Status int         `json:"status"`         // 状态码
	Msg    string      `json:"msg,omitempty"`  // 其他信息
	Data   T           `json:"data,omitempty"` // 响应数据
	Page   *ResultPage `json:"page,omitempty"` // 分页信息
}

// ResultPage 列表分页数据
type ResultPage struct {
	PageIndex int `json:"pageIndex"` // 当前页码
	PageSize  int `json:"pageSize"`  // 每页数量
	PageCount int `json:"pageCount"` // 总页数
	DataCount int `json:"dataCount"` // 数据总量
}

func (r Result[T]) ToAny() Result[any] {
	return Result[any]{
		Status: r.Status,
		Msg:    r.Msg,
		Data:   r.Data,
		Page:   r.Page,
	}
}

// Success 成功响应
func Success[T any](data T) Result[T] {
	return Result[T]{
		Status: http.StatusOK,
		Data:   data,
	}
}

// SuccessWithPage 成功带分页响应
func SuccessWithPage[T any](data T, pageIndex, pageSize, pageCount, dataCount int) Result[T] {
	page := &ResultPage{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		PageCount: pageCount,
		DataCount: dataCount,
	}

	return Result[T]{
		Status: http.StatusOK,
		Data:   data,
		Page:   page,
	}
}

// Error 失败响应
func Error[T any](status int, msg string) Result[T] {
	return Result[T]{
		Status: status,
		Msg:    msg,
	}
}
