package routers

import (
	"blogserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/", &controllers.HomeController{}, "get:Index")
	//beego.Router("/2", &controllers.HomeController{}, "get:Index2")
	//beego.Router("/explore", &controllers.ExploreController{}, "*:Index")

	//登录、注册、修改用户资料
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Register")

	// 编辑、发布文章
	beego.Router("/editArticle", &controllers.MarkdownStore{}, "post:Edit")
	beego.Router("/editLanguage", &controllers.Language{}, "post:Edit")
}
