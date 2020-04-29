package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"encoding/json"
	"fmt"
)

type MarkdownStore struct {
	BaseController
}

func (m *MarkdownStore) Edit() {
	store := models.NewMarkdownStore()
	data := m.Ctx.Input.RequestBody
	json.Unmarshal(data, store) //解析二进制json，把结果放进ob中
	//obj := &models.MarkdownStore{UserId: store.UserId,LanguageId: store.LanguageId,Content: store.Content}
	err := store.Edit("languageTitle", "htmlContent")
	if nil != err {
		fmt.Println(err)
		m.Data["json"] = common.ResultHandle(map[string]bool{"result": true}, err)
	} else {
		m.Data["json"] = common.ResultHandle(map[string]interface{}{"result": true, "msg": "新增成功"}, err)
	}
	m.ServeJSON()
}
