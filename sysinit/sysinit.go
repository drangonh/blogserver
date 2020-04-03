/**
 * @Author: dragon
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2020/4/2 下午5:45
 */

package sysinit

import (
	"github.com/astaxie/beego"
	"path/filepath"
	"strings"
)

func sysInit() {
	uploads := filepath.Join("/", "uploads")
	beego.BConfig.WebConfig.StaticDir["uploads"] = uploads

	registerFunctions()
}

func registerFunctions() {
	beego.AddFuncMap("cdnjs", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdnjs", "")
		if strings.HasPrefix(p, "/") && strings.HasPrefix(cdn, "/") {
			return cdn + string(p[1:])
		}

		return cdn
	})
}
