package model

// OperationLogResp 操作日志响应结果
type OperationLogResp struct {
	ObjectType   string `json:"objectType"`   // 对象类型
	ObjectId     string `json:"objectId"`     // 对象ID
	Method       string `json:"method"`       // 操作方式
	OperatorId   string `json:"operatorId"`   // 操作人ID
	OperatorName string `json:"operatorName"` // 操作人名
	CreatedAt    int64  `json:"createdAt"`    // 创建时间戳
}

// UserResp 用户响应结果
type UserResp struct {
	UserId    string `json:"userId"`    // 用户ID
	UserName  string `json:"userName"`  // 用户名
	LoginIp   string `json:"loginIp"`   // 登录IP
	LoginAt   int64  `json:"loginAt"`   // 登录时间戳
	CreatedAt int64  `json:"createdAt"` // 创建时间戳
	UpdatedAt int64  `json:"updatedAt"` // 更新时间戳
}

// CategoryResp 文章分类响应结果
type CategoryResp struct {
	CateId      string `json:"cateId"`      // 分类ID
	Name        string `json:"name"`        // 名称
	Description string `json:"description"` // 描述
	State       int    `json:"state"`       // 状态
	Sort        int    `json:"sort"`        // 排序
	ParentId    string `json:"parentId"`    // 父类ID
	CreatedAt   int64  `json:"createdAt"`   // 创建时间戳
	UpdatedAt   int64  `json:"updatedAt"`   // 更新时间戳
}

// ArticleResp 文章响应结果
type ArticleResp struct {
	CateId      string `json:"cateId"`      // 分类ID
	ArticleId   string `json:"articleId"`   // 文章ID
	Title       string `json:"title"`       // 标题
	Description string `json:"description"` // 描述
	State       int    `json:"state"`       // 状态
	Sort        int    `json:"sort"`        // 排序
	PageView    int    `json:"pageView"`    // 访问量
	CreatedAt   int64  `json:"createdAt"`   // 创建时间戳
	UpdatedAt   int64  `json:"updatedAt"`   // 更新时间戳
}
