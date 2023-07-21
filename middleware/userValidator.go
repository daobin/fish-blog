package middleware

import "github.com/daobin/goeasy"

func CheckLogin() func(c *goeasy.Context) {
	return func(c *goeasy.Context) {
		// todo 上线使用时需要完善此登录校验逻辑（系统当前主要是为了校验框架goeasy，暂时没有上线使用）
	}
}
