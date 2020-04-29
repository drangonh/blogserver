package models

import "github.com/astaxie/beego/orm"

type MarkdownStoreModel struct {
	ContentId   int    `orm:"pk;auto;column(contentId)" json:"contentId"`
	UserId      int    ` orm:"column(userId)" json:"userId"`
	LanguageId  int    `orm:"type(int);column(languageId)" json:"languageId"`               //对应语音的ID
	Content     string `orm:"type(text);column(content)" json:"content"`                    //文章内容对应的是markdown的字符串
	htmlContent string `orm:"type(text);column(htmlContent)" json:"htmlContent" default:""` //文章内容：markdown内容对应的html，暂时不会存值
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

	if one.ContentId > 0 {
		_, err = o.Update(c, str...)
	} else {
		_, err = o.Insert(c)
	}

	return
}
