package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type MarkdownStoreModel struct {
	ContentId   int    `orm:"pk;auto;column(contentId)" json:"contentId"`
	UserId      int    ` orm:"column(userId)" json:"userId"`
	LanguageId  int    `orm:"type(int);column(languageId)" json:"languageId"`    //对应语音的ID
	Content     string `orm:"type(text);column(content)" json:"content"`         //文章内容对应的是markdown的字符串
	StoreTitle  string `orm:"type(text);column(storeTitle)" json:"storeTitle"`   //文章标题
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

	// 直接截取字符串可能出现截取了字符串中一个字符的部分字节，截取字符串最后的部分乱码。最终
	// 导致出现程序异常
	// 解决办法是把string转成[]rune(str),最后转成string保存即可
	newStr := []rune(c.Content)
	if len(newStr) < 100 {
		c.Brief = c.Content
	} else {
		arr := newStr[0:100]
		c.Brief = string(arr)
	}

	fmt.Println("简介：：：", c.Brief)
	if one.ContentId > 0 {
		_, err = o.Update(c, str...)
		fmt.Println(err)

	} else {
		_, err = o.Insert(c)
	}

	return
}

func (c *MarkdownStoreModel) GetList(userId int, languageId int, isBrief bool) (list []MarkdownStoreModel, err error) {
	o := orm.NewOrm()

	// languageId,查询用户所有的文章
	if languageId == 0 {
		//查询简介
		if isBrief {
			_, err = o.Raw("SELECT contentId, userId, storeTitle, languageId, brief, htmlContent FROM "+
				TNMarkdownStore()+" WHERE userId = ?", userId).QueryRows(&list)
		} else {
			// 查询详情
			_, err = o.Raw("SELECT contentId, userId, storeTitle, languageId, content, htmlContent FROM "+
				TNMarkdownStore()+" WHERE userId = ?", userId).QueryRows(&list)
		}
	} else {
		//查询简介
		if isBrief {
			_, err = o.Raw("SELECT contentId, userId, storeTitle, languageId, brief, htmlContent FROM "+
				TNMarkdownStore()+" WHERE userId = ? AND languageId = ?", userId, languageId).QueryRows(&list)
		} else {
			// 查询详情
			_, err = o.Raw("SELECT contentId, userId, storeTitle, languageId, content, htmlContent FROM "+
				TNMarkdownStore()+" WHERE userId = ? AND languageId = ?", userId, languageId).QueryRows(&list)
		}
	}

	return
}

// 获取详情
func (c *MarkdownStoreModel) GetDetail(userId int, languageId int, contentId int) (detail MarkdownStoreModel, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(TNMarkdownStore()).Filter("userId", userId).Filter("languageId", languageId).Filter("contentId", contentId).One(&detail)
	return
}

// 删除文章
func (c *MarkdownStoreModel) DeletePage(contentId int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(TNMarkdownStore()).Filter("contentId", contentId).Delete()
	return
}
