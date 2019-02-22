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

func GiveRoleGetByUsername(username string) (*GameName,error){
	g := new(GameName)
	err := orm.NewOrm().QueryTable(TableName("giverole")).Filter("user_name",username).One(g)
	if err != nil{
		return nil,err
	}
	return g,nil
}