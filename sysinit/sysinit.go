/**
 * @Author: dragon
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2020/4/2 下午5:45
 */

package sysinit

import (
	_ "blogserver/models"
	"blogserver/utils"
	"github.com/astaxie/beego"
	"path/filepath"
	"strings"
)

func sysInit() {

	//设置上传的静态资源路径，此处就是当前目录下的uploads
	uploads := filepath.Join("./", "uploads")
	beego.BConfig.WebConfig.StaticDir["uploads"] = uploads

	//设置文档资源路径
	beego.SetStaticPath("/swagger", "swagger")
	//if beego.BConfig.RunMode == "dev" {
	//	beego.SetStaticPath("/swagger", "swagger")
	//}

	//注册views中需要调用的后端的函数
	registerFunctions()
}

func registerFunctions() {

	//cdnjs这个函数就是返回cdnjs的全路径
	beego.AddFuncMap("cdnjs", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdnjs", "")
		if strings.HasPrefix(p, "/") && strings.HasPrefix(cdn, "/") {
			return cdn + string(p[1:])
		}

		return cdn + p
	})

	beego.AddFuncMap("cdncss", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdncss", "")
		if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
			return cdn + string(p[1:])
		}
		if !strings.HasPrefix(p, "/") && !strings.HasSuffix(cdn, "/") {
			return cdn + "/" + p
		}
		return cdn + p
	})

	//beego.AddFuncMap("getUsernameByUid", func(id interface{}) string {
	//	return new(models.Member).GetUsernameByUid(id)
	//})
	//beego.AddFuncMap("getNicknameByUid", func(id interface{}) string {
	//	return new(models.Member).GetNicknameByUid(id)
	//})
	beego.AddFuncMap("inMap", utils.InMap)

	//	//用户是否收藏了文档
	//beego.AddFuncMap("doesCollection", new(models.Collection).DoesCollection)
	//	beego.AddFuncMap("scoreFloat", utils.ScoreFloat)
	beego.AddFuncMap("showImg", utils.ShowImg)
	//beego.AddFuncMap("IsFollow", new(models.Fans).Relation)
	beego.AddFuncMap("isubstr", utils.Substr)
}
