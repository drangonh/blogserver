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
		new(Category),
		new(Book),
		new(BookCategory),
	)
}

/*
* Table Names
* */

func TNCategory() string {
	return "md_category"
}

func TNBookCategory() string {
	return "md_book_category"
}

func TNBook() string {
	return "md_books"
}
