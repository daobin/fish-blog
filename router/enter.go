package router

import (
	"github.com/daobin/fish-blog/conf"
	"github.com/daobin/fish-blog/controller/admin"
	_ "github.com/daobin/fish-blog/docs"
	"github.com/daobin/fish-blog/middleware"
	"github.com/daobin/goeasy"
	"net/http"
)

func InitRouter(engine *goeasy.Engine) {
	// 开发环境展示Swag文档接口
	if conf.App.GetString("env") == "dev" {
		engine.GET("/docs/*any", func(c *goeasy.Context) {
			// todo 开发环境展示Swag文档接口
		})
	}

	// 路由中间件：初始化数据库配置
	engine.Use(middleware.InitDbConf())

	// 后台路由初始化
	initAdminRouter(engine)
	// 前台路由初始化
	initIndexRouter(engine)
}

func initAdminRouter(engine *goeasy.Engine) {
	// 路由分组
	rtGroup := engine.Group("admin")
	// 路由中间件：校验用户登录状态
	rtGroup.Use(middleware.CheckLogin())

	// 分类
	rtGroup.POST("category", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Category.Add(c))
	})
	rtGroup.PUT("category", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Category.Update(c))
	})
	rtGroup.DELETE("category", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Category.Delete(c))
	})
	rtGroup.GET("category", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Category.GetInfo(c))
	})
	rtGroup.GET("category/list", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Category.GetList(c))
	})
	rtGroup.GET("category/tree", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Category.GetTree(c))
	})

	// 文章
	rtGroup.POST("article", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Article.Add(c))
	})
	rtGroup.PUT("article", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Article.Update(c))
	})
	rtGroup.GET("article/:id", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Article.GetInfo(c))
	})
	rtGroup.GET("article/list", func(c *goeasy.Context) {
		c.Json(http.StatusOK, admin.Article.GetList(c))
	})

}

func initIndexRouter(engine *goeasy.Engine) {
	// todo 注册前台路由
}
