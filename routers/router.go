package routers

import (
	"blogserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// web入口
	beego.Router("/", &controllers.HomeController{}, "*:Index")

	//登录、注册、修改用户资料
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Register")

	// 文章
	beego.Router("/editArticle", &controllers.MarkdownStore{}, "post:Edit")
	beego.Router("/getArticleList", &controllers.MarkdownStore{}, "get:GetMarkdownList")
	beego.Router("/getArticleDetail", &controllers.MarkdownStore{}, "get:GetDetail")

	// 语言
	beego.Router("/getLanguageList", &controllers.Language{}, "get:GetLanguageList")
	beego.Router("/editLanguage", &controllers.Language{}, "post:Edit")
}
