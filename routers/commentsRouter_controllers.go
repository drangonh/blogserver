package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blogserver/controllers:HomeController"] = append(beego.GlobalControllerRouter["blogserver/controllers:HomeController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/`,
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:Language"] = append(beego.GlobalControllerRouter["blogserver/controllers:Language"],
		beego.ControllerComments{
			Method:           "Edit",
			Router:           `/editLanguage`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:Language"] = append(beego.GlobalControllerRouter["blogserver/controllers:Language"],
		beego.ControllerComments{
			Method:           "GetLanguageList",
			Router:           `/getLanguageList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

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

	beego.GlobalControllerRouter["blogserver/controllers:OSSController"] = append(beego.GlobalControllerRouter["blogserver/controllers:OSSController"],
		beego.ControllerComments{
			Method:           "EditInfo",
			Router:           `/editOss`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:OSSController"] = append(beego.GlobalControllerRouter["blogserver/controllers:OSSController"],
		beego.ControllerComments{
			Method:           "OssInfo",
			Router:           `/ossInfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["blogserver/controllers:ProfileController"] = append(beego.GlobalControllerRouter["blogserver/controllers:ProfileController"],
		beego.ControllerComments{
			Method:           "EditUserProfile",
			Router:           `/editUserProfile`,
			AllowHTTPMethods: []string{"post"},
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
