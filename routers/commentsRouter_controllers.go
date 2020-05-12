package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blogserver/controllers:MarkdownStore"] = append(beego.GlobalControllerRouter["blogserver/controllers:MarkdownStore"],
		beego.ControllerComments{
			Method:           "DeletePage",
			Router:           `/deleteArticle`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:MarkdownStore"] = append(beego.GlobalControllerRouter["blogserver/controllers:MarkdownStore"],
		beego.ControllerComments{
			Method:           "Edit",
			Router:           `/editArticle`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:MarkdownStore"] = append(beego.GlobalControllerRouter["blogserver/controllers:MarkdownStore"],
		beego.ControllerComments{
			Method:           "GetDetail",
			Router:           `/getArticleDetail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:MarkdownStore"] = append(beego.GlobalControllerRouter["blogserver/controllers:MarkdownStore"],
		beego.ControllerComments{
			Method:           "GetMarkdownList",
			Router:           `/getArticleList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:UserController"] = append(beego.GlobalControllerRouter["blogserver/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:UserController"] = append(beego.GlobalControllerRouter["blogserver/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:UserController"] = append(beego.GlobalControllerRouter["blogserver/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Register",
			Router:           `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
