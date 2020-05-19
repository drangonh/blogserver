package models

import (
	common2 "blogserver/common"
	"errors"
	"github.com/astaxie/beego/orm"
	"regexp"
)

//用户信息详情表
type Profile struct {
	Id          int    `orm:"pk;auto" json:"id"`
	Avatar      string `orm:"size(30),unique;column(avatar)" json:"avatar"`
	Email       string `orm:"size(100);unique" json:"email"`
	Description string `orm:"column(description)" json:"description"`
	User        *User  `orm:"rel(one)" json:"userId"`
}

//orm 回调
func (m *Profile) TableName() string {
	return TNProfile()
}

func NewProfile() *Profile {
	return &Profile{}
}

// 新增、修改用户信息
func (m *Profile) EditProfile(strS ...string) (err error) {
	if ok, err := regexp.MatchString(common2.RegexpEmail, m.Email); !ok || err != nil || m.Email == "" {
		return errors.New("邮箱格式错误")
	}

	o := orm.NewOrm()

	// 查询数据库是否存在该条数据,外键在数据库中会自动加"_id"
	p := &Profile{User: m.User}
	o.Read(p, "user_id")

	if err != nil {
		return
	}

	if p.Id > 0 {
		//修改用户信息
		sql := "UPDATE user_profile SET avatar = ?, email = ?, description = ? WHERE id = ?"
		_, err = o.Raw(sql, m.Avatar, m.Email, m.Description, p.Id).Exec()
	} else {
		_, err = o.Insert(m)
	}
	return
}
