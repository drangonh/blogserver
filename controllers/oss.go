package controllers

import (
	"blogserver/utils/common"
	"blogserver/utils/oss"
	"encoding/json"
)

type OSSController struct {
	BaseController
}

type UpLoad struct {
	fileName string
	path     string
	fileType int
}

// @Title UploadImage
// @Description 图片上传
// @Param fileName formData string true "文件名"
// @Param path formData string true "文件路径"
// @Param fileType formData int true "文件类型"
// formData 表示是 post 请求的数据，query 表示带在 url 之后的参数，path 表示请求路径上得参数，
// 例如上面例子里面的 key，body 表示是一个 raw 数据请求，header 表示带在 header 信息中得参数。
// @Success 200 true
// @Failure nil
// @router /uploadImage [post]
func (o *OSSController) UploadImage() {
	obj := &UpLoad{}
	data := o.Ctx.Input.RequestBody
	json.Unmarshal(data, &obj)

	err := oss.UploadFile(obj.fileName, obj.path, obj.fileType)
	if err != nil {
		o.Data["json"] = common.ResultHandle(nil, err)
	} else {
		o.Data["json"] = common.ResultHandle(true, nil)
	}
}
