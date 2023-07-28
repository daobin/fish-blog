package model

import "gopkg.in/mgo.v2/bson"

// OperationLogEntity 操作日志 [mongo]
type OperationLogEntity struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	ObjectType   string        `json:"objectType" bson:"objectType"`     // 对象类型
	ObjectId     string        `json:"objectId" bson:"objectId"`         // 对象ID
	Method       string        `json:"method" bson:"method"`             // 操作方式
	OperatorId   string        `json:"operatorId" bson:"operatorId"`     // 操作人ID
	OperatorName string        `json:"operatorName" bson:"operatorName"` // 操作人名
	CreatedAt    int64         `json:"createdAt" bson:"createdAt"`       // 创建时间戳
}

// UserEntity 用户 [mongo]
type UserEntity struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	UserId    string        `json:"userId" bson:"userId"`       // 用户ID
	UserName  string        `json:"userName" bson:"userName"`   // 用户名
	Password  string        `json:"password" bson:"password"`   // 登录密码
	LoginIp   string        `json:"loginIp" bson:"loginIp"`     // 登录IP
	LoginAt   int64         `json:"loginAt" bson:"loginAt"`     // 登录时间戳
	CreatedAt int64         `json:"createdAt" bson:"createdAt"` // 创建时间戳
	UpdatedAt int64         `json:"updatedAt" bson:"updatedAt"` // 更新时间戳
}

// CategoryEntity 分类 [mongo]
type CategoryEntity struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	CateId      string        `json:"cateId" bson:"cateId"`           // 分类ID
	Name        string        `json:"name" bson:"name"`               // 名称
	Description string        `json:"description" bson:"description"` // 描述
	State       int           `json:"state" bson:"state"`             // 状态
	Sort        int           `json:"sort" bson:"sort"`               // 排序
	ParentId    string        `json:"parentId" bson:"parentId"`       // 父类ID
	CreatedAt   int64         `json:"createdAt" bson:"createdAt"`     // 创建时间戳
	UpdatedAt   int64         `json:"updatedAt" bson:"updatedAt"`     // 更新时间戳
}

// ArticleEntity 文章 [mongo]
type ArticleEntity struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	CateId      string        `json:"cateId" bson:"cateId"`           // 分类ID
	ArticleId   string        `json:"articleId" bson:"articleId"`     // 文章ID
	Title       string        `json:"title" bson:"title"`             // 标题
	Description string        `json:"description" bson:"description"` // 描述
	State       int           `json:"state" bson:"state"`             // 状态
	Sort        int           `json:"sort" bson:"sort"`               // 排序
	PageView    int           `json:"pageView" bson:"pageView"`       // 访问量
	CreatedAt   int64         `json:"createdAt" bson:"createdAt"`     // 创建时间戳
	UpdatedAt   int64         `json:"updatedAt" bson:"updatedAt"`     // 更新时间戳
}
