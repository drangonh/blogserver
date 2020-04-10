package models

type Member struct {
	MemberId int `orm:"pk,auto" json:"member_id"`
}

//orm 回调
func (m *Member) TableName() string {
	return TNMember()
}

func NewMember() *Member {
	return &Member{}
}
