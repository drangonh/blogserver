package models

import (
	"blogserver/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

func ElasticSearchDocument(kw string, pageSize, page int) (ids []int, count int, err error) {
	if page > 0 {
		page = page - 1
	} else {
		page = 0
	}
	//搜索全部
	queryJson := `
		{
		    "query" : {
		        "match" : {
		        	"content":"%v"
		        }
		    },
		    "_source":["contentId"],
			"size": %v,
			"from": %v
		}
	`
	queryJson = fmt.Sprintf(queryJson, kw, pageSize, page)

	//elasticsearch api
	host := beego.AppConfig.String("elastic_host")
	api := host + "mdocuments/datas/_search"

	fmt.Println(api)
	fmt.Println(queryJson)

	sj, err := utils.HttpPostJson(api, queryJson)
	if nil == err {
		count = sj.GetPath("hits", "total").MustInt()
		resultArray := sj.GetPath("hits", "hits").MustArray()

		fmt.Println("ES数据获取：：：：", resultArray)
		for _, v := range resultArray {
			if each_map, ok := v.(map[string]interface{}); ok {
				id, _ := strconv.Atoi(each_map["_id"].(string))
				ids = append(ids, id)
			}
		}
	}
	return ids, count, err
}

// 一次更新一个分类的文章到ES
func ElasticBuildIndex(userId string, languageId string) {
	var documents []MarkdownStoreModel
	fields := []string{"contentId", "userId", "languageId", "content"}

	//筛选分类中所有的文章
	GetOrm("w").QueryTable(TNMarkdownStore()).Filter("userId", userId).Filter("languageId", languageId).All(&documents, fields...)
	if len(documents) > 0 {
		for _, document := range documents {
			////更新ES 中索引 mdocuments 的内容
			AddDocumentToIndex(document.ContentId, document.UserId, document.LanguageId, document.Content)
		}
	}
}

//把文章存入ES中
func AddDocumentToIndex(contentId, userId, languageId int, content string) {
	queryJson := `
		{
			"contentId":%v,
			"userId":%v,
			"languageId":%v,
			"content":"%v"
		}
	`

	//elasticsearch api
	host := beego.AppConfig.String("elastic_host")
	api := host + "mdocuments/datas/" + strconv.Itoa(contentId)

	//发起请求
	queryJson = fmt.Sprintf(queryJson, contentId, userId, languageId, strings.Replace(content, "\n", "", -1))
	err := utils.HttpPutJson(api, queryJson)
	//fmt.Println(api)
	//fmt.Println(queryJson)
	//fmt.Println("ES数据新增成功", err)
	if nil != err {
		beego.Debug(err)
	}

}
