package controllers

import "fmt"

type HomeController struct {
	BaseController
}

func (u *HomeController) URLMapping() {
	u.Mapping("Index", u.Index)
}

// @Title Index
// @Description 网站入口
// @Success 200 true
// @Failure nil
// @router / [*]
func (u *HomeController) Index() {
	fmt.Println("网页进入")

	u.TplName = "index.html"
}
