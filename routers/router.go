package routers

import (
	"blogserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//登录、注册、修改用户资料
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Register")

	// 文章
	beego.Router("/getLanguageList", &controllers.Language{}, "get:GetLanguageList")
	beego.Router("/editArticle", &controllers.MarkdownStore{}, "post:Edit")

	// 语言
	beego.Router("/editLanguage", &controllers.Language{}, "post:Edit")
}
