package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Role struct {
	Id int
	Role string
	RoleName string
	RoleLevel string
	HigherRole string
	Roleid int
	RolePid int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);null"`
	//GameName *GameName `orm:"rel(fk)"`
}


func (r *Role) TableName() string{
	return TableName("role")
}

func RoleAdd(r *Role) (int64,error) {
	return orm.NewOrm().Insert(r)
}


func RoleGetByroleName(rolename string) (*Role, error) {
	r := new(Role)
	err := orm.NewOrm().QueryTable(TableName("role")).Filter("role",rolename).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func RoleGetByHigherRole(higherrole string) (*Role,error){
	r := new(Role)
	err := orm.NewOrm().QueryTable(TableName("role")).Filter("role",higherrole).One(r)
	if err != nil{
		return nil,err
	}
	return r,nil
}

func (r *Role) RoleUpdate(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}

func (r *Role) RoleDeleteByRolename(rolename string) (int64,error){
	query := orm.NewOrm().QueryTable(TableName("role"))
	num,err := query.Filter("role",rolename).Delete()
	return num,err
}

func (r *Role) RoleDeleteByRoleid (roleid int) (int64,error){
	query := orm.NewOrm().QueryTable(TableName("role"))
	num,err := query.Filter("roleid",roleid).Delete()
	return num,err
}