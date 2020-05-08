package controllers

import (
	common2 "blogserver/common"
	"blogserver/models"
	"blogserver/utils"
	"blogserver/utils/common"
	"encoding/json"
	"fmt"
	"time"
)

type UserController struct {
	BaseController
}

type LoginModal struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

func (u *UserController) Index2() string {
	return "测试一下"
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [Post]
func (u *UserController) Login() {
	var remember CookieRemember
	var loginObj LoginModal
	data := u.Ctx.Input.RequestBody
	json.Unmarshal(data, &loginObj)
	fmt.Println(loginObj, data)
	if user, err := models.NewUser().Login(loginObj.UserName, loginObj.PassWord); err == nil {
		u.Data["json"] = common.ResultHandle(user, nil)

		//登录成功之后设置加密的cookie
		remember.UserId = user.UserId
		remember.Time = time.Now()
		v, _ := utils.Encode(remember)

		//登录成功设置cookie
		u.SetSecureCookie(common2.AppKey(), "login", v, common2.CookieMastLiftTime)
	} else {
		fmt.Println("登录错误", err)
		u.Data["json"] = common.ResultHandle(nil, err)
	}

	u.ServeJSON()
}

//退出登录,清除登录cookie和session
func (u *UserController) Logout() {
	u.SetMember(models.User{})
	u.SetSecureCookie(common2.AppKey(), "login", "", -3600)

	obj := map[string]interface{}{
		"res": true,
		"msg": "退出登录",
	}
	u.Data["json"] = common.ResultHandle(obj, nil)
	u.ServeJSON()
}

func (u *UserController) Register() {
	var addUser models.AddUser
	data := u.Ctx.Input.RequestBody
	json.Unmarshal(data, &addUser) //解析二进制json，把结果放进ob中
	user := &models.User{Username: addUser.Username, Password: addUser.Password}
	res, err := user.AddUser(addUser) //这是添加用户函数
	if nil != err {
		fmt.Println(err)

		u.Data["json"] = common.ResultHandle(nil, err)
	} else {
		u.Data["json"] = common.ResultHandle(res, nil)
	}
	u.ServeJSON()
}
