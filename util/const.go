package util

const SortDefault = 50

// 状态
const (
	StateEnable  = iota + 1 // 启用
	StateDisable            // 禁用
	StateDelete             // 删除
)

// 对象类型
const (
	ObjTypeCate    = "category" // 分类
	ObjTypeArticle = "article"  // 文章
	ObjTypeUser    = "user"     // 用户
)
