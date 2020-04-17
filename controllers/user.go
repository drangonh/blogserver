package controllers

import (
	"blogserver/models"
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
	username := u.GetString("username")
	password := u.GetString("password")
	if user, err := models.NewUser().Login(username, password); err == nil {
		u.Data["json"] = user
	} else {
		u.Data["json"] = err
	}
	u.ServeJSON()
}

func (u *UserController) Register() {
	var addUser models.AddUser
	data := u.Ctx.Input.RequestBody
	json.Unmarshal(data, &addUser) //解析二进制json，把结果放进ob中
	user := &models.User{Username: addUser.Username, Password: addUser.Password}
	err := user.AddUser(addUser) //这是添加用户函数
	if nil != err {
		fmt.Println(err)
		u.Data["json"] = map[string]interface{}{"result": false, "msg": err}
	} else {
		u.Data["json"] = map[string]interface{}{"result": true, "msg": "新增成功"}
	}
	u.ServeJSON()
}
