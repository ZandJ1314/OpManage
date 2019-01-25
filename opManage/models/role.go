package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Role struct {
	Id int
	RoleName string
	RoleDescription string
	RoleLevel string
	HigherRole string
	RolePid int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);null"`
	Gamename *GameName `orm:"rel(fk)"`
}

func (r *Role) TableName() string{
	return TableName("role")
}

func RoleAdd(r *Role) (int64,error) {
	return orm.NewOrm().Insert(r)
}


func RoleGetById(id int) (*Role, error) {
	r := new(Role)
	err := orm.NewOrm().QueryTable(TableName("role")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Role) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}