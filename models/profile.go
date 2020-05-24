package models

import (
	common2 "blogserver/common"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"regexp"
)

//用户信息详情表
type Profile struct {
	Id          int    `orm:"pk;auto" json:"id"`
	Avatar      string `orm:"size(30);column(avatar)" json:"avatar"`
	Email       string `orm:"size(100);unique" json:"email"`
	Description string `orm:"column(description)" json:"description"`
	NickName    string `orm:"size(30);unique;column(nickName)" json:"nickName"`
	Uid         int    `orm:"column(uid);unique" json:"uid"`
}

//orm 回调
func (m *Profile) TableName() string {
	return TNProfile()
}

func NewProfile() *Profile {
	return &Profile{}
}

// 新增、修改用户信息
func (m *Profile) EditProfile(strS ...string) (info *Profile, err error) {
	if ok, err := regexp.MatchString(common2.RegexpEmail, m.Email); !ok || err != nil || m.Email == "" {

		fmt.Println(err, m.Email)
		return nil, errors.New("邮箱格式错误")
	}

	cond := orm.NewCondition().Or("email", m.Email).Or("nickName", m.NickName).Or("uid", m.Uid)
	one := NewProfile()
	o := GetOrm("w")
	err = o.QueryTable(m.TableName()).SetCond(cond).One(one)

	if one.NickName == m.NickName {
		return nil, errors.New("用户名已被使用")
	}

	if one.Email == m.Email {
		return nil, errors.New("邮箱已被使用")
	}

	if one.Id > 0 {
		//修改用户信息
		m.Id = one.Id
		_, err = o.Update(m, strS...)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = o.Insert(m)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
