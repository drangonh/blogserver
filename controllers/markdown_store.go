package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"encoding/json"
	"errors"
	"fmt"
)

type MarkdownStore struct {
	BaseController
}

type deleteModal struct {
	contentId int `json:"contentId"`
}

// 新增或者插入新的文章
func (m *MarkdownStore) Edit() {
	store := models.NewMarkdownStore()
	data := m.Ctx.Input.RequestBody
	json.Unmarshal(data, store) //解析二进制json，把结果放进ob中

	//obj := &models.MarkdownStore{UserId: store.UserId,LanguageId: store.LanguageId,Content: store.Content}
	err := store.Edit("content", "htmlContent", "storeTitle", "brief")
	if nil != err {
		m.Data["json"] = common.ResultHandle(map[string]bool{"result": true}, err)
	} else {
		m.Data["json"] = common.ResultHandle(map[string]interface{}{"result": true, "msg": "新增成功"}, err)
	}
	m.ServeJSON()
}

//查询指定用户指定分类的文章,也可以查询指定客户所有的文章
func (m *MarkdownStore) GetMarkdownList() {
	userId, _ := m.GetInt("userId")

	// isBrief
	// true查询简介和标题
	// false查询详情和标题
	// 默认查询简介

	isBrief, _ := m.GetBool("isBrief", true)

	if userId == 0 {
		m.Data["json"] = common.ResultHandle(nil, errors.New("请求参数userId错误"))
	} else {
		languageId, _ := m.GetInt("languageId")

		List, err := models.NewMarkdownStore().GetList(userId, languageId, isBrief)

		if err != nil {
			m.Data["json"] = common.ResultHandle(nil, err)
		} else {
			m.Data["json"] = common.ResultHandle(List, nil)
		}
	}

	m.ServeJSON()
}

// 获取详情
func (m *MarkdownStore) GetDetail() {
	contentId, _ := m.GetInt("contentId")
	userId, _ := m.GetInt("userId")
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

// 删除文章
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
