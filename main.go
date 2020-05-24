package main

import (
	_ "blogserver/routers"
	_ "blogserver/sysinit"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	//InsertFilter是提供一个过滤函数，后端处理跨越问题
	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	//允许访问所有源
	//	AllowAllOrigins: true,
	//	//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
	//	//其中Options跨域复杂请求预检
	//	AllowMethods: []string{"*"},
	//	//指的是允许的Header的种类
	//	AllowHeaders: []string{"*"},
	//	//公开的HTTP标头列表
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	//如果设置，则允许共享身份验证凭据，例如cookie
	//	AllowCredentials: true,
	//}))

	// 这段代码放在router.go文件的init()的开头
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Authorization",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type"},
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type"},
		AllowCredentials: true,
	}))

	beego.Run()

}
