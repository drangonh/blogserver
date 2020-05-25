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
	UserName string `orm:"column(userName)" json:"userName"`
	PassWord string `orm:"column(passWord)" json:"passWord"`
}

func (u *UserController) URLMapping() {
	u.Mapping("Login", u.Login)
	u.Mapping("Logout", u.Logout)
	u.Mapping("Register", u.Register)
}

// @Title Login
// @Description Logs user into the system
// @Param	userName		formData 	string	true		"The username for login"
// @Param	passWord		formData 	string	true		"The password for login"
// @Success 200 {userName,password,userId} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var remember CookieRemember
	var loginObj LoginModal
	data := u.Ctx.Input.RequestBody
	json.Unmarshal(data, &loginObj)
	fmt.Println(loginObj, data)
	if user, err := models.NewUser().Login(loginObj.UserName, loginObj.PassWord); err == nil {
		u.Data["json"] = common.ResultHandle(user, nil)

		//登录成功之后设置加密的cookie
		remember.UserId = user.Uid
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

// @Title Logout
// @Description 退出登录,清除登录cookie和session
// @Success 200 {"res": true,"msg": "退出登录",}
// @router /logout [get]
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

// @Title Register
// @Description 账号注册
// @Param Username formData string true "The Username for register"
// @Param Password formData string true "The Password for register"
// @Param ConfirmPassword formData string true "The ConfirmPassword for register"
// @Success 200 {userName,password,userId}
// @Failure 注册失败
// @router /register [post]
func (u *UserController) Register() {
	var addUser models.AddUser
	data := u.Ctx.Input.RequestBody
	json.Unmarshal(data, &addUser) //解析二进制json，把结果放进ob中
	user := &models.User{Username: addUser.Username, Password: addUser.Password, LastLoginTime: time.Now()}

	fmt.Println(addUser)

	res, err := user.AddUser(addUser) //这是添加用户函数
	if nil != err {
		fmt.Println(err)

		u.Data["json"] = common.ResultHandle(nil, err)
	} else {
		u.Data["json"] = common.ResultHandle(res, nil)
	}
	u.ServeJSON()
}

// @Title GetArtAndCatCnt
// @Description 根据用户id获取用户的分类总数和文章总数
// @Success 200 true
// @Failure nil
// @router /getArtAndCatCnt [get]
func (u *UserController) GetArtAndCatCnt() {
	uid := u.User.UserId

	art, cat, err := models.NewUser().FindArtCatCnt(uid)

	u.Data["json"] = common.ResultHandle(map[string]int{
		"language": cat,
		"article":  art,
	}, err)

	u.ServeJSON()
}
