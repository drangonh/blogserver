/**
 * @Author: dragon
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2020/4/3 上午10:05
 */

package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(

		new(Profile),
		new(LanguageModel),
		new(MarkdownStoreModel),
		new(User),
	)
}

/*
* Table Names
* */

// 用户表
func TNUser() string {
	return "user"
}

// 用户信息表
func TNProfile() string {
	return "user_profile"
}

// 计算机开发语言表
func TNLanguage() string {
	return "language"
}

// 文章表
func TNMarkdownStore() string {
	return "markdown_store"
}

/*
* Tool Funcs
* */
//获取orm对象
//@param alias 数据库alias
func GetOrm(alias string) orm.Ormer {
	o := orm.NewOrm()
	if len(alias) > 0 {
		beego.Debug("Using Alias : " + alias)
		if "w" == alias {
			o.Using("default")
		} else {
			o.Using(alias)
		}
	}

	return o
}
