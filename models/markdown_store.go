package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
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

// b.contentId,b.userId,b.languageId,b.content,b.storeTitle,l.languageTitle
type Detail struct {
	ContentId     int    `orm:"column(contentId)" json:"contentId"`
	UserId        int    `orm:"column(userId)" json:"userId"`
	LanguageId    int    `orm:"column(languageId)" json:"languageId"`
	Content       string `orm:"column(content)" json:"content"`
	StoreTitle    string `orm:"column(storeTitle)" json:"storeTitle"`
	LanguageTitle string `orm:"column(languageTitle)" json:"languageTitle"`
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
	} else {
		_, err = o.Insert(c)
	}

	return
}

//分页查询第一步：查询总数和对应分页的contentId数组
func (c *MarkdownStoreModel) GetListCount(userId int, languageId int, page int, size int) (list []MarkdownStoreModel, cnt int, err error) {
	var count struct {
		Cnt int
	}

	o := orm.NewOrm()
	fields := "contentId, userId"
	// languageId,查询用户所有的文章
	var sql, sqlCount string
	if languageId == 0 {
		sql = "SELECT %v  FROM " +
			TNMarkdownStore() + " WHERE userId = ?"
		//查询总数
		sqlCount = fmt.Sprintf(sql, "count(contentId) cnt")
		//查询该分页的id
		sql = fmt.Sprintf(sql, fields)
		err = o.Raw(sqlCount, userId).QueryRow(&count)

	} else {
		sql = "SELECT %v FROM " +
			TNMarkdownStore() + " WHERE userId = ? AND languageId = ?"
		sqlCount = fmt.Sprintf(sql, "count(contentId) cnt")
		sql = fmt.Sprintf(sql, fields)
		err = o.Raw(sqlCount, userId, languageId).QueryRow(&count)
	}

	cnt = count.Cnt

	if cnt <= 0 {
		return
	}

	limit := fmt.Sprintf(" limit %v offset %v", size, (page-1)*size)

	sql = sql + limit

	if languageId == 0 {
		_, err = o.Raw(sql, userId).QueryRows(&list)
	} else {
		_, err = o.Raw(sql, userId, languageId).QueryRows(&list)
	}
	return
}

func (c *MarkdownStoreModel) GetList(ids []int, isBrief bool) (list []MarkdownStoreModel, err error) {
	o := orm.NewOrm()

	var strS []string

	for _, item := range ids {
		strS = append(strS, strconv.Itoa(item))
	}

	//查询简介
	if isBrief {
		sql := "SELECT contentId, userId, storeTitle, languageId, brief, htmlContent FROM " +
			TNMarkdownStore() + " WHERE contentId in (%v)"
		sql = fmt.Sprintf(sql, strings.Join(strS, ","))
		_, err = o.Raw(sql).QueryRows(&list)
	} else {
		// 查询详情
		sql := "SELECT contentId, userId, storeTitle, languageId, content, htmlContent FROM " +
			TNMarkdownStore() + " WHERE contentId in (%v)"
		sql = fmt.Sprintf(sql, strings.Join(strS, ","))
		_, err = o.Raw(sql).QueryRows(&list)
	}

	return
}

// 获取详情
func (c *MarkdownStoreModel) GetDetail(userId int, languageId int, contentId int) (detail Detail, err error) {
	var rows []Detail
	o := GetOrm("w")
	sql := "select b.contentId,b.userId,b.languageId,b.content,b.storeTitle,l.languageTitle " +
		"from markdown_store as b left join language as l on b.languageId = l.languageId " +
		"where b.userId = ? and b.languageId = ? and b.contentId = ?;"

	num, err := o.Raw(sql, userId, languageId, contentId).QueryRows(&rows)

	fmt.Println(rows, num)

	if len(rows) > 0 {
		detail = rows[0]
	}
	return
}

// 获取上一篇文章或者下一篇文章的详情
// 获取详情
func (c *MarkdownStoreModel) GetLastOrNextDetail(userId int, languageId int, contentId int, artType int) (detail Detail, err error) {
	var rows []Detail
	o := GetOrm("w")
	var sqlLast string
	if artType == 1 {
		sqlLast = "select * from markdown_store as m left join language as l on m.languageId = l.languageId " +
			"where m.contentId < ? and m.userId = ? and m.languageId = ? order by m.contentId desc limit 1;"
	} else {
		sqlLast = "select * from markdown_store as m left join language as l on m.languageId = l.languageId " +
			"where m.contentId > ? and m.userId = ? and m.languageId = ? order by m.contentId asc limit 1;"
	}

	num, err := o.Raw(sqlLast, contentId, userId, languageId).QueryRows(&rows)

	fmt.Println(rows, num)

	if len(rows) > 0 {
		detail = rows[0]
	}
	return
}

// 删除文章
func (c *MarkdownStoreModel) DeletePage(contentId int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(TNMarkdownStore()).Filter("contentId", contentId).Delete()
	return
}
