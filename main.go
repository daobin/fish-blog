package main

import (
	"github.com/daobin/fish-blog/router"
	"github.com/daobin/goeasy"
)

//@Title Fish-Blog
//@Version 1.0.0
//@Description Fish个人博客

func main() {
	// 新建框架引擎
	engine := goeasy.New()

	// 初始化路由
	router.InitRouter(engine)

	// 启动框架引擎
	goeasy.Start(":80")
}
