package model

import (
	"errors"
	"time"
)

var _ = time.Thursday

//User
type User struct {
	Uid        uint   `gorm:"column:uid" form:"uid" json:"uid" comment:"" columnType:"int(10) unsigned" dataType:"int" columnKey:"PRI"`
	Name       string `gorm:"column:name" form:"name" json:"name" comment:"" columnType:"varchar(32)" dataType:"varchar" columnKey:"UNI"`
	Password   string `gorm:"column:password" form:"password" json:"password" comment:"" columnType:"varchar(64)" dataType:"varchar" columnKey:""`
	Mail       string `gorm:"column:mail" form:"mail" json:"mail" comment:"" columnType:"varchar(150)" dataType:"varchar" columnKey:"UNI"`
	Url        string `gorm:"column:url" form:"url" json:"url" comment:"" columnType:"varchar(150)" dataType:"varchar" columnKey:""`
	Screenname string `gorm:"column:screenName" form:"screenName" json:"screenName" comment:"" columnType:"varchar(32)" dataType:"varchar" columnKey:""`
	Created    uint   `gorm:"column:created" form:"created" json:"created" comment:"" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
	Activated  uint   `gorm:"column:activated" form:"activated" json:"activated" comment:"" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
	Logged     uint   `gorm:"column:logged" form:"logged" json:"logged" comment:"" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
	Group      string `gorm:"column:group" form:"group" json:"group" comment:"" columnType:"varchar(16)" dataType:"varchar" columnKey:""`
	Authcode   string `gorm:"column:authCode" form:"authCode" json:"authCode" comment:"" columnType:"varchar(64)" dataType:"varchar" columnKey:""`
}

//TableName
func (m *User) TableName() string {
	return "users"
}

//One
func (m *User) One() (one *User, err error) {
	one = &User{}
	err = crudOne(m, one)
	return
}

//All
func (m *User) All(q *PaginationQuery) (list *[]User, total uint, err error) {
	list = &[]User{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *User) Update() (err error) {
	where := User{Uid: m.Uid}
	m.Uid = 0

	return crudUpdate(m, where)
}

//CreateUserOfRole
func (m *User) CreateUserOfRole() error {
	m.Uid = 0

	return db.Create(m).Error
}

//Delete
func (m *User) Delete() error {
	if m.Uid == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
