/**
 * @Author: dragon
 * @Description:
 * @File:  testorm
 * @Version: 1.0.0
 * @Date: 2020/4/2 下午3:39
 */

package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
	Id   int
	Name string
	Sex  int8
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:yn666666@tcp(localhost:3306)/blogserver?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "小红红", Sex: 1}
	o.Insert(&user)
}
