package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type LanguageModel struct {
	LanguageId      int    `orm:"pk;auto;column(languageId)" json:"languageId"`
	UserId          int    `orm:"column(userId)" json:"userId"`
	LanguageTitle   string `orm:"type(text);column(languageTitle);unique" json:"languageTitle"` //语言标题
	LanguageContent string `orm:"type(text);column(languageContent)" json:"languageContent"`    //语言简洁
}

func (c *LanguageModel) TableName() string {
	return TNLanguage()
}

func NewLanguage() *LanguageModel {
	return &LanguageModel{}
}

func (c *LanguageModel) Edit(str ...string) (err error) {
	o := orm.NewOrm()
	var one LanguageModel
	o.QueryTable(TNLanguage()).Filter("languageId", c.LanguageId).One(&one)

	if one.LanguageId > 0 {
		_, err = o.Update(c, str...)
	} else {
		_, err = o.Insert(c)
	}
	return
}

func (c *LanguageModel) GetList(id string) (list []LanguageModel, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(TNLanguage()).Filter("userId", id).One(&list)

	fmt.Println(list)
	return
}
