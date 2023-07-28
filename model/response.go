package model

// CategoryTreeResp 分类树响应结果
type CategoryTreeResp struct {
	CategoryEntity
	Children []CategoryTreeResp `json:"children"`
}
