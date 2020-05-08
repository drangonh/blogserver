package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"encoding/json"
	"fmt"
)

type Language struct {
	BaseController
}

func (m *Language) Edit() {
	fmt.Println("获取baseController中获取到的用户信息", m.User)
	language := models.NewLanguage()
	data := m.Ctx.Input.RequestBody
	json.Unmarshal(data, language) //解析二进制json，把结果放进ob中
	//obj := &models.MarkdownStore{UserId: store.UserId,LanguageId: store.LanguageId,Content: store.Content}
	err := language.Edit("languageContent", "languageTitle")
	if nil != err {
		fmt.Println(err)
		m.Data["json"] = common.ResultHandle(map[string]bool{"result": true}, err)
	} else {
		m.Data["json"] = common.ResultHandle(map[string]interface{}{"result": true, "msg": "新增成功"}, err)
	}
	m.ServeJSON()
}

func (m *Language) GetLanguageList() {
	fmt.Println("获取baseController中获取到的用户信息", m.User)
	language := models.NewLanguage()
	userId := m.GetString("userId")
	list, err := language.GetList(userId)

	if err != nil {
		m.Data["json"] = common.ResultHandle(nil, err)
	} else {
		m.Data["json"] = common.ResultHandle(list, nil)
	}

	m.ServeJSON()
}
