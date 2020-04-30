package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"encoding/json"
	"errors"
)

type MarkdownStore struct {
	BaseController
}

// 新增或者插入新的文章
func (m *MarkdownStore) Edit() {
	store := models.NewMarkdownStore()
	data := m.Ctx.Input.RequestBody
	json.Unmarshal(data, store) //解析二进制json，把结果放进ob中

	//obj := &models.MarkdownStore{UserId: store.UserId,LanguageId: store.LanguageId,Content: store.Content}
	err := store.Edit("content", "htmlContent")
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
	if userId == 0 {
		m.Data["json"] = common.ResultHandle(nil, errors.New("请求参数userId错误"))
	} else {
		languageId, _ := m.GetInt("languageId")

		List, err := models.NewMarkdownStore().GetList(userId, languageId)

		if err != nil {
			m.Data["json"] = common.ResultHandle(nil, err)
		} else {
			m.Data["json"] = common.ResultHandle(List, nil)
		}
	}

	m.ServeJSON()
}