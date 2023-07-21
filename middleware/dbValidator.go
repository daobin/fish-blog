package middleware

import (
	"github.com/daobin/fish-blog/conf/db"
	"github.com/daobin/fish-blog/model"
	"github.com/daobin/fish-blog/util"
	"github.com/daobin/goeasy"
	"net/http"
)

// InitDbConf 初始化数据库配置
func InitDbConf() func(c *goeasy.Context) {
	return func(c *goeasy.Context) {
		// 初始化Mongo
		err := db.Mongo.InitFile()
		if err != nil {
			c.Json(http.StatusInternalServerError, model.Error[any](util.ErrCodeSystem.ToInt(), util.ErrCodeSystem.GetText()))
			c.Abort()
			return
		}
	}
}
