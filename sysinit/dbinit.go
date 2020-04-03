/**
 * @Author: dragon
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2020/4/2 下午5:45
 */

package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func dbInit(alias ...string) {
	if len(alias) > 0 {
		for _, alia := range alias {
			registerDatabase(alia)

			if alia == "w" {
				orm.RunSyncdb("default", false, true)
			}
		}
	} else {
		registerDatabase("w")
		orm.RunSyncdb("default", false, true)
	}

	//设置是否为开发模式，控制打印
	isDev := beego.AppConfig.String("runmode") == "dev"

	if isDev {
		//设置开发模式下可以打印日志
		orm.Debug = isDev
	}
}

func registerDatabase(alias string) {
	dbAlias := alias
	if alias == "w" || alias == "default" || len(alias) <= 0 {
		dbAlias = "default"
		alias = "w"
	}

	//数据库名
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	//数据库用户
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	//数据库密码
	dbPwd := beego.AppConfig.String("db_" + alias + "_password")
	//数据库IP
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	//数据库端口
	dbPort := beego.AppConfig.String("db_" + alias + "_port")

	//	orm.RegisterDataBase("default", "mysql", "root:yn666666@tcp(localhost:3306)/blogserver?charset=utf8", 30)
	dataSource := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	fmt.Println(dataSource)
	orm.RegisterDataBase(dbAlias, "mysql", dataSource, 30)
}
