package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"encoding/json"
	"fmt"
)

type ProfileController struct {
	BaseController
}

func (u *ProfileController) URLMapping() {
	u.Mapping("EditUserProfile", u.EditUserProfile)
}

// @Title EditUserProfile
// @Description 更新或者修改用户信息
// @Param avatar formData string true "用户头像"
// @Param email formData string true "用户邮箱"
// @Param nickName formData string true "用户昵称"
// @Param description formData string true "用户描述"
// @Success 200 {userName,password,userId}
// @Failure 注册失败
// @router /editUserProfile [post]
func (u *ProfileController) EditUserProfile() {
	var profile models.Profile
	data := u.Ctx.Input.RequestBody
	json.Unmarshal(data, &profile)

	fmt.Println(profile)
	info, err := profile.EditProfile("avatar", "email", "description", "nickName")

	if nil != err {
		u.Data["json"] = common.ResultHandle(nil, err)
	} else {
		u.Data["json"] = common.ResultHandle(info, nil)
	}

	u.ServeJSON()
}
