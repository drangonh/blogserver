package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Language struct {
	BaseController
}

func (m *Language) URLMapping() {
	m.Mapping("Edit", m.Edit)
	m.Mapping("GetLanguageList", m.GetLanguageList)
}

// @Title Edit
// @Description 语言编辑或者新增
// @Param languageTitle formData  string true "语言标题"
// @Param languageContent formData string true "语言简介"
// @Success 200 true
// @Failure nil
// @router /editLanguage [post]
func (m *Language) Edit() {
	language := models.NewLanguage()
	data := m.Ctx.Input.RequestBody
	json.Unmarshal(data, language) //解析二进制json，把结果放进ob中
	//obj := &models.MarkdownStore{UserId: store.UserId,LanguageId: store.LanguageId,Content: store.Content}

	if language.LanguageTitle == "" {
		m.Data["json"] = common.ResultHandle(map[string]bool{"result": false}, errors.New("语言标题不能为空"))
		m.ServeJSON()
		return
	}

	if language.LanguageContent == "" {
		m.Data["json"] = common.ResultHandle(map[string]bool{"result": false}, errors.New("语言简介不能为空"))
		m.ServeJSON()
		return
	}

	language.UserId = m.User.UserId

	id, err := language.Edit("languageContent", "languageTitle")
	if nil != err {
		fmt.Println(err)
		m.Data["json"] = common.ResultHandle(map[string]bool{"result": true}, err)
	} else {
		fmt.Println("文章ID：：：：", id)
		go func() {
			models.ElasticBuildIndex(string(m.User.UserId), string(id))
		}()
		m.Data["json"] = common.ResultHandle(map[string]interface{}{"result": true, "msg": "新增成功"}, err)
	}
	m.ServeJSON()
}

// @Title Edit
// @Description 语言列表
// @Success 200 list
// @Failure nil
// @router /getLanguageList [get]
func (m *Language) GetLanguageList() {
	language := models.NewLanguage()
	id := m.User.UserId
	list, err := language.GetList(strconv.Itoa(id))

	if err != nil {
		m.Data["json"] = common.ResultHandle(nil, err)
	} else {
		m.Data["json"] = common.ResultHandle(list, nil)
	}

	m.ServeJSON()
}
