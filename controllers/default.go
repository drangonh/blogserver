package controllers

import (
	"blogserver/models"
	"fmt"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Hi() {
	fmt.Println("hello go")
	c.Data["Website"] = "beego.me000"
	c.TplName = "index.tpl"
	models.PrintUsers()
}
