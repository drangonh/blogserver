package controllers

import (
	"blogserver/models"
	"blogserver/utils/common"
	"errors"
)

type ElasticSearchController struct {
	BaseController
}

// @Title SearchResult
// @Description 全文搜索
// @Param wd query int true "文章ID"
// @Param page query int true "文章第几页"
// @Param isBrief query bool true "查询简介的列表还是详情的列表，默认查询简介"
// @Success 200 {"list":  List,"count": count}
// @Failure nil
// @router /searchDocuments [get]
func (c *ElasticSearchController) SearchResult() {
	//获取关键词
	wd := c.GetString("wd")
	if "" == wd {
		c.Data["json"] = common.ResultHandle(nil, errors.New("请输入搜索关键字"))
		c.ServeJSON()
		return
	}
	isBrief, _ := c.GetBool("isBrief", true)

	//page&size
	page, _ := c.GetInt("page", 1)
	if page < 1 {
		page = 1
	}
	size := 10

	// 获取分页的slice id和总共的文章数量
	ids, count, err := models.ElasticSearchDocument(wd, size, page)
	if nil == err && len(ids) > 0 {
		list, err := models.NewMarkdownStore().GetList(ids, isBrief)

		if err != nil {
			c.Data["json"] = common.ResultHandle(nil, err)
		} else {
			c.Data["json"] = common.ResultHandle(
				map[string]interface{}{
					"list":  list,  //分页的列表
					"count": count, //文章总数
				}, nil)
		}
	} else {
		c.Data["json"] = common.ResultHandle(nil, err)
	}

	c.ServeJSON()

}
