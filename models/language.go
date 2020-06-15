package models

import (
	"github.com/astaxie/beego/orm"
)

type LanguageModel struct {
	LanguageId      int    `orm:"pk;auto;column(languageId)" json:"languageId"`
	Pid             int    `orm:"column(pid)" json:"pid"` //父级id，实现无限极分类
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

func (c *LanguageModel) Edit(str ...string) (id int64, err error) {
	o := orm.NewOrm()
	var one LanguageModel
	o.QueryTable(TNLanguage()).Filter("languageId", c.LanguageId).One(&one)

	if one.LanguageId > 0 {
		id, err = o.Update(c, str...)
	} else {
		id, err = o.Insert(c)
	}
	return
}

func (c *LanguageModel) GetList(id string) (list []LanguageModel, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(TNLanguage()).Filter("userId", id).All(&list)

	return
}

func (c *LanguageModel) GetOne(id string, languageId string) (detail LanguageModel, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(TNLanguage()).Filter("userId", id).Filter("languageId", languageId).One(&detail)

	return
}
