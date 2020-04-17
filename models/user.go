package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strings"
)

type User struct {
	UserId   int    `orm:"pk;auto" json:"userId"`
	Username string `orm:"size(30) unique" json:"userName"`
	Password string `orm:"size(50) unique" json:"passWord"`
}

type AddUser struct {
	UserId          int    `orm:"pk;auto" json:"userId"`
	Username        string `orm:"size(30) unique" json:"userName"`
	Password        string `orm:"size(50) unique" json:"passWord"`
	ConfirmPassword string `orm:"size(50) unique" json:"confirmPassWord"`
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
	err := orm.NewOrm().QueryTable(m.TableName()).Filter("userName", userName).One(member)

	if err != nil {
		return member, errors.New("用户不存在")
	}

	if member.Password == password {
		return member, nil
	}
	return member, errors.New("密码错误")
}

//注册账号
func (m *User) AddUser(add AddUser) error {
	//验证信息
	if add.Password != add.ConfirmPassword {
		return errors.New("两次输入的密码请确保一致")
	}

	if l := strings.Count(add.Password, ""); l < 6 || l >= 20 {
		return errors.New("密码请输入6-20个字符")
	}

	//start,查询账号是否存在
	o := orm.NewOrm()
	var user User
	o.QueryTable(m.TableName()).Filter("userName", m.Username).One(&user)
	if user.Username == m.Username {
		return errors.New("用户名已存在")
	}
	//end

	//插入数据
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}
