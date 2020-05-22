package models

import (
	"blogserver/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type User struct {
	UserId        int       `orm:"pk;auto;column(userId)" json:"userId"`
	Username      string    `orm:"size(30),unique;column(userName)" json:"userName"`
	Password      string    `orm:"size(50),unique;column(passWord)" json:"passWord"`
	LastLoginTime time.Time `orm:"type(datetime);column(lastLoginTime);null" json:"lastLoginTime"`
	//LastLoginTime time.Time `orm:"type(datetime);column(lastLoginTime)" json:"lastLoginTime"` //如果不是
	//Profile       *Profile  `orm:"reverse(one)"` //// Reverse relationship (optional)
}

type AddUser struct {
	UserId          int    `json:"userId"`
	Username        string `json:"userName"`
	Password        string `json:"passWord"`
	ConfirmPassword string `json:"confirmPassWord"`
}

//orm 回调
func (m *User) TableName() string {
	return TNUser()
}

func NewUser() *User {
	return &User{}
}

func (m *User) Login(userName, password string) (*Profile, error) {
	o := GetOrm("w")
	member := &User{}
	err := o.QueryTable(m.TableName()).Filter("userName", userName).One(member)

	member.LastLoginTime = time.Now().Local()

	// 更新登录时间
	_, err = o.Update(member, "lastLoginTime")
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, errors.New("用户不存在")
	}

	res, err := utils.PasswordVerify(member.Password, password)

	if res == false {
		return nil, errors.New("密码错误")
	}

	var profile *Profile

	o.QueryTable(TNProfile()).Filter("uid", member.UserId).One(profile)
	return profile, nil
}

//注册账号
func (m *User) AddUser(add AddUser) (*User, error) {
	//验证信息
	if add.Password != add.ConfirmPassword {
		return nil, errors.New("两次输入的密码请确保一致")
	}

	fmt.Println(add)
	if l := strings.Count(add.Password, ""); l < 6 || l >= 20 {
		return nil, errors.New("请输入6-20位密码")
	}

	//start,查询账号是否存在
	o := orm.NewOrm()
	var user User

	o.QueryTable(m.TableName()).Filter("userName", m.Username).One(&user)
	if user.Username == m.Username {
		return nil, errors.New("用户名已存在")
	}
	//end

	hash, err := utils.PasswordHash(m.Password)
	if err == nil {
		m.Password = hash
	}

	//插入数据
	_, err = o.Insert(m)
	if err != nil {
		return nil, err
	}

	fmt.Println("新增用户", user)
	return m, nil
}

func (m *User) Find(id int) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{UserId: id}

	fmt.Println("查询用户信息：", user)
	err = o.Read(user)
	return user, err
}
