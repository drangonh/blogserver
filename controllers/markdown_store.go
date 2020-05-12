package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
)

type MarkdownStore struct {
	BaseController
}

type deleteModal struct {
	contentId int `json:"contentId"`
}

// 文章
//beego.Router("/editArticle", &controllers.MarkdownStore{}, "post:Edit")
//beego.Router("/getArticleList", &controllers.MarkdownStore{}, "get:GetMarkdownList")
//beego.Router("/getArticleDetail", &controllers.MarkdownStore{}, "get:GetDetail")
//beego.Router("/deleteArticle", &controllers.MarkdownStore{}, "post:DeletePage")

// 注解接口
func (m *MarkdownStore) URLMapping() {
	m.Mapping("Edit", m.Edit)
	m.Mapping("GetMarkdownList", m.GetMarkdownList)
	m.Mapping("GetDetail", m.GetDetail)
	m.Mapping("DeletePage", m.DeletePage)
}

// @Title Edit
// @Description 新增或者插入新的文章
// @Param languageId query int "文章ID"
// @Param content query string "文章内容"
// @Param storeTitle query string "文章标题"
// @Param htmlContent query string "文章内容对应的html"
// @Success 200 {"result": true, "msg": "新增成功"}
// @Failure {"result": false}
// @router /editArticle [post]
func (m *MarkdownStore) Edit() {
	store := models.NewMarkdownStore()
	data := m.Ctx.Input.RequestBody
	json.Unmarshal(data, store) //解析二进制json，把结果放进ob中
	store.UserId = m.User.UserId
	err := store.Edit("content", "htmlContent", "storeTitle", "brief", "languageId")
	if nil != err {
		m.Data["json"] = common.ResultHandle(map[string]bool{"result": false}, err)
	} else {
		m.Data["json"] = common.ResultHandle(map[string]interface{}{"result": true, "msg": "新增成功"}, err)
	}
	m.ServeJSON()
}

// @Title GetMarkdownList
// @Description 查询指定用户指定分类的文章,也可以查询指定客户所有的文章
// @Param languageId query int "文章ID"
// @Param page query int "文章第几页"
// @Param pageSize query int "一个分页文章的文章个数"
// @Param isBrief query bool "查询简介的列表还是详情的列表，默认查询简介"
// @Success 200 {"list":  List,"count": count}
// @Failure nil
// @router /getArticleList [get]
func (m *MarkdownStore) GetMarkdownList() {
	var contentIds []int
	id := m.User.UserId
	// isBrief
	// true查询简介和标题
	// false查询详情和标题
	// 默认查询简介
	isBrief, _ := m.GetBool("isBrief", true)

	page, err := m.GetInt("page", 1)
	if err != nil {
		m.Data["json"] = common.ResultHandle(nil, errors.New("请求参数page缺少"))
	}
	size, err := m.GetInt("pageSize", 1)
	if err != nil {
		m.Data["json"] = common.ResultHandle(nil, errors.New("请求参数pageSize缺少"))
	}
	if id == 0 {
		m.Data["json"] = common.ResultHandle(nil, errors.New("请求参数userId错误"))
	} else {
		languageId, _ := m.GetInt("languageId")

		ids, count, err := models.NewMarkdownStore().GetListCount(id, languageId, page, size)

		fmt.Println("分页的数组", ids)
		if err != nil {
			beego.Error(err.Error())
		} else {
			for _, obj := range ids {
				contentIds = append(contentIds, obj.ContentId)
			}
		}

		List, err := models.NewMarkdownStore().GetList(contentIds, isBrief)

		if err != nil {
			m.Data["json"] = common.ResultHandle(nil, err)
		} else {
			m.Data["json"] = common.ResultHandle(
				map[string]interface{}{
					"list":  List,  //分页的列表
					"count": count, //文章总数
				}, nil)
		}
	}

	m.ServeJSON()
}

// @Title GetDetail
// @Description 获取详情
// @Param languageId query int "文章ID"
// @Param contentId query int "文章id"
// @Success 200 {"list":  List,"count": count}
// @Failure nil
// @router /getArticleDetail [get]
func (m *MarkdownStore) GetDetail() {
	contentId, _ := m.GetInt("contentId")
	userId := m.User.UserId
	languageId, _ := m.GetInt("languageId")

	fmt.Println(contentId, userId, languageId)
	detail, err := models.NewMarkdownStore().GetDetail(userId, languageId, contentId)

	if err != nil {
		m.Data["json"] = common.ResultHandle(nil, err)
	} else {
		m.Data["json"] = common.ResultHandle(detail, nil)
	}
	m.ServeJSON()
}

// @Title DeletePage
// @Description 删除文章
// @Param contentId query int "文章id"
// @Success 200 true
// @Failure nil
// @router /deleteArticle [post]
func (m *MarkdownStore) DeletePage() {
	var del map[string]int
	obj := m.Ctx.Input.RequestBody
	json.Unmarshal(obj, &del)

	fmt.Println(del, obj)
	err := models.NewMarkdownStore().DeletePage(del["contentId"])

	if err != nil {
		m.Data["json"] = common.ResultHandle(false, err)
	} else {
		m.Data["json"] = common.ResultHandle(true, nil)
	}
	m.ServeJSON()
}
