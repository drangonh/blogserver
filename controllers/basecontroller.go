/**
 * @Author: dragon
 * @Description:
 * @File:  basecontroller
 * @Version: 1.0.0
 * @Date: 2020/4/3 上午10:06
 */

package controllers

import (
	"blogserver/common"
	"blogserver/models"
	"blogserver/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type BaseController struct {
	beego.Controller
	User *models.User
}

type CookieRemember struct {
	UserId int
	Time   time.Time
}

//每个子类Controller公用方法调用前，都执行一下Prepare方法
func (c *BaseController) Prepare() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//CORS(Cross-Origin Resource Sharing，跨域资源共享)允许设置，浏览器现在都自动支持了CORS
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	c.User = models.NewUser() //初始化
	//从session中获取用户信息
	if user, ok := c.GetSession(common.SessionName).(models.User); ok && user.UserId > 0 {
		fmt.Println("获取session值")
		c.User = &user
	} else {
		//如果Cookie中存在登录信息，从cookie中获取用户信息
		if cookie, ok := c.GetSecureCookie(common.AppKey(), "login"); ok {
			fmt.Println("获取cookie值")
			var remember CookieRemember
			err := utils.Decode(cookie, &remember)

			if err == nil {
				user, err := models.NewUser().Find(remember.UserId)
				if err == nil {
					c.SetMember(*user)
					c.User = user
				}
			}
		}
	}
}

// 设置登录用户信息
func (c *BaseController) SetMember(user models.User) {
	if user.UserId <= 0 {
		c.DelSession(common.SessionName)
		c.DelSession("uid")
		c.DestroySession()
	} else {
		c.SetSession(common.SessionName, user)
		c.SetSession("uid", user.UserId)
	}
}
