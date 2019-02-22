package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserType struct {
	Id int
	Usertypename string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);null"`
	//User *User `orm:"rel(fk)"`  //设置一对多的关系
	//设置一对多的关系
	//User []*User `orm:"reverse(many)"`
}

func (u *UserType) TableName() string{
	return TableName("usertype")
}

func UserTypeAdd(u *UserType) (int64,error){
	return orm.NewOrm().Insert(u)
}

func UserTypegetByuserTypename(name string) (*UserType,error){
	t := new(UserType)
	err := orm.NewOrm().QueryTable(TableName("usertype")).Filter("usertypename",name).One(t)
	if err != nil{
		return nil,err
	}
	return t,err
}
