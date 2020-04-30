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
		new(LanguageModel),
		new(MarkdownStoreModel),
	)
}

/*
* Table Names
* */

// 用户表
func TNUser() string {
	return "user"
}

// 计算机开发语言表
func TNLanguage() string {
	return "language"
}

// 文章表
func TNMarkdownStore() string {
	return "markdown_store"
}
