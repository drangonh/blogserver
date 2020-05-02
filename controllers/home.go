package controllers

type HomeController struct {
	BaseController
}

func (u *HomeController) Index() {
	u.TplName = "index.html"
}
