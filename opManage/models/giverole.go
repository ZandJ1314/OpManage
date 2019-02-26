package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type GiveRole struct {
	Id int
	UserName string
	GameName string
	Role string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);null"`
}

func (r *GiveRole) TableName() string{
	return TableName("giverole")
}

func GiveRoleAdd(g *GiveRole) (int64,error) {
	return orm.NewOrm().Insert(g)
}

func GiveRoleGetByUsername(username string) (*GiveRole,error){
	g := new(GiveRole)
	err := orm.NewOrm().QueryTable(TableName("giverole")).Filter("user_name",username).One(g)
	if err != nil{
		return nil,err
	}
	return g,nil
}

func (g *GiveRole) GiveRoleUpdate(fields ...string) error {
	if _, err := orm.NewOrm().Update(g, fields...); err != nil {
		return err
	}
	return nil
}