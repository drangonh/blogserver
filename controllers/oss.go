package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"encoding/json"
)

type OSSController struct {
	BaseController
}

// @Title OssInfo
// @Description 获取oss信息
// @Success 200 true
// @Failure nil
// @router /ossInfo [get]
func (o *OSSController) OssInfo() {
	oss := models.NewOSS()
	info, err := oss.GetOSSInfo()
	if err != nil {
		o.Data["json"] = common.ResultHandle(nil, err)
	} else {
		o.Data["json"] = common.ResultHandle(info, nil)
	}
	o.ServeJSON()
}

// @Title EditInfo
// @Description 新增或者编辑oss信息
// @Param uid formData int true "用户ID"
// @Param endpoint formData string true "Your region:类似：oss-cn-hangzhou"
// @Param accessKeyId formData string true "Your AccessKeyId"
// @Param accessKeySecret formData string true "Your AccessKeySecret"
// @Param bucketName formData string true "Your bucket name"
// @Success 200 true
// @Failure nil
// @router /editOss [post]
func (o *OSSController) EditInfo() {
	oss := models.NewOSS()
	data := o.Ctx.Input.RequestBody
	json.Unmarshal(data, oss)

	err := oss.Edit()
	if err != nil {
		o.Data["json"] = common.ResultHandle(nil, err)
	} else {
		o.Data["json"] = common.ResultHandle(true, nil)
	}
	o.ServeJSON()
}
