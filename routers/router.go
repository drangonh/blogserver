// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"blogserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/blog",
		beego.NSInclude(
			&controllers.HomeController{},
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/language",
			beego.NSInclude(
				&controllers.Language{},
			),
		),
		beego.NSNamespace("/markdownStore",
			beego.NSInclude(
				&controllers.MarkdownStore{},
			),
		),
	)
	beego.AddNamespace(ns)
}
