package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

type User struct {
	UserId   int    `orm:"pk,auto" json:"userId"`
	Username string `orm:"size(30) unique" json:"userName"`
	Password string `orm:"size(50) unique" json:"passWord"`
}

//orm 回调
func (m *User) TableName() string {
	return TNUser()
}

func NewUser() *User {
	return &User{}
}

func (m *User) Login(userName, password string) (*User, error) {
	member := &User{}
	err := orm.NewOrm().QueryTable(m.TableName()).Filter("userName", userName).Filter("status", 0).One(member)

	if err != nil {
		return member, errors.New("用户不存在")
	}

	if member.Password == password {
		return member, nil
	}
	return member, errors.New("密码错误")
}
