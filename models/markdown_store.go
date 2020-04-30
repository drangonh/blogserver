package models

import (
	"github.com/astaxie/beego/orm"
)

type MarkdownStoreModel struct {
	ContentId   int    `orm:"pk;auto;column(contentId)" json:"contentId"`
	UserId      int    ` orm:"column(userId)" json:"userId"`
	LanguageId  int    `orm:"type(int);column(languageId)" json:"languageId"`    //对应语音的ID
	Content     string `orm:"type(text);column(content)" json:"content"`         //文章内容对应的是markdown的字符串
	Brief       string `orm:"type(text);column(brief)" json:"brief"`             //文章内容对应的是markdown的字符串前100个字符
	HtmlContent string `orm:"type(text);column(htmlContent)" json:"htmlContent"` //文章内容：markdown内容对应的html，暂时不会存值
}

func (c *MarkdownStoreModel) TableName() string {
	return TNMarkdownStore()
}

func NewMarkdownStore() *MarkdownStoreModel {
	return &MarkdownStoreModel{}
}

func (c *MarkdownStoreModel) Edit(str ...string) (err error) {
	o := orm.NewOrm()
	var one MarkdownStoreModel
	o.QueryTable(TNMarkdownStore()).Filter("contentId", c.ContentId).One(&one)

	if len(c.Content) < 100 {
		c.Brief = c.Content
	} else {
		c.Brief = c.Content[0:100]
	}

	if one.ContentId > 0 {
		_, err = o.Update(c, str...)
	} else {
		_, err = o.Insert(c)
	}

	return
}

func (c *MarkdownStoreModel) GetList(userId int, languageId int, isBrief bool) (list *[]MarkdownStoreModel, err error) {
	o := orm.NewOrm()

	if languageId == 0 {
		//查询简介
		if isBrief {
			_, err = o.Raw("SELECT contentId, userId, languageId, brief, htmlContent FROM "+
				TNMarkdownStore()+" WHERE userId = ? AND isBrief", userId, true).QueryRows(list)
		} else {
			// 查询详情
			_, err = o.Raw("SELECT contentId, userId, languageId, content, htmlContent FROM "+
				TNMarkdownStore()+" WHERE userId = ? AND isBrief", userId, true).QueryRows(list)
		}
	} else {
		//查询简介
		if isBrief {
			_, err = o.Raw("SELECT contentId, userId, languageId, brief, htmlContent FROM "+
				TNMarkdownStore()+" WHERE userId = ? AND languageId = ?", userId, languageId).QueryRows(list)
		} else {
			// 查询详情
			_, err = o.Raw("SELECT contentId, userId, languageId, content, htmlContent FROM "+
				TNMarkdownStore()+" WHERE userId = ? AND languageId = ?", userId, languageId).QueryRows(list)
		}
	}

	return
}
