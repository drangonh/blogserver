package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"encoding/json"
	"fmt"
)

type UserController struct {
	BaseController
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
// @router /login [get]
func (u *UserController) Login() {
	userName := u.GetString("userName")
	passWord := u.GetString("passWord")
	if user, err := models.NewUser().Login(userName, passWord); err == nil {
		u.Data["json"] = common.ResultHandle(user, nil)
	} else {
		fmt.Println("登录错误", err)
		u.Data["json"] = common.ResultHandle(nil, err)
	}
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
