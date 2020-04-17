/**
 * @Author: dragon
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2020/4/3 上午10:05
 */

package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(
		new(User),
	)
}

/*
* Table Names
* */

func TNUser() string {
	return "user"
}
