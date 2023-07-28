package model

type ObjectIdReq struct {
	Id string `json:"id"` // 操作对象ID
}

// SaveUserReq 保存用户请求
type SaveUserReq struct {
	UserName string `json:"userName"` // 用户名
	Password string `json:"password"` // 登录密码
}

// SaveCategoryReq 保存分类请求
type SaveCategoryReq struct {
	CateId      string `json:"cateId"`      // 分类ID
	Name        string `json:"name"`        // 名称
	Description string `json:"description"` // 描述
	State       int    `json:"state"`       // 状态
	Sort        int    `json:"sort"`        // 排序
	ParentId    string `json:"parentId"`    // 父类ID
}

// GetCategoryListReq 获取分类列表请求
type GetCategoryListReq struct {
	Name         string `json:"name"`         // 名称
	ParentId     string `json:"parentId"`     // 父类ID
	State        int    `json:"state"`        // 状态
	IsReturnPage bool   `json:"isReturnPage"` // 是否返回分页，默认：否
	PageIndex    int    `json:"pageIndex"`    // 查询页码，默认：1
	PageSize     int    `json:"pageSize"`     // 每页数量，默认：10
}

// SaveArticleReq 保存文章请求
type SaveArticleReq struct {
	CateId      string `json:"cateId"`      // 分类ID
	ArticleId   string `json:"articleId"`   // 文章ID
	Title       string `json:"title"`       // 标题
	Description string `json:"description"` // 描述
	State       int    `json:"state"`       // 状态
	Sort        int    `json:"sort"`        // 排序
	PageView    int    `json:"pageView"`    // 访问量
}

// GetArticleListReq 获取文章列表请求
type GetArticleListReq struct {
	CateIds      string `json:"cateIds"`      // 分类ID，多个之间用逗号分隔
	State        int    `json:"state"`        // 状态
	IsReturnPage bool   `json:"isReturnPage"` // 是否返回分页，默认：否
	PageIndex    int    `json:"pageIndex"`    // 查询页码，默认：1
	PageSize     int    `json:"pageSize"`     // 每页数量，默认：10
}
