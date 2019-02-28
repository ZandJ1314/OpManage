package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	Openid string `orm:"null"`
	UserName string
	AlaisUserName string
	Department string
	PhoneNumber string `orm:"null"`
	EmailAdress string `orm:"null"`
	HeadPortraitName string `orm:"null"`
	ManageName string
	Issuperadministrator int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`  //第一次保存时才设置时间
	UpdateTime time.Time `orm:"auto_now;type(datetime);null"`   //每一次宝宝都会对时间进行更新
	//UserType *UserType `orm:"rel(fk)"`//设置一对多反向关系
	//Gametype []*GameType `orm:"reverse(many)"`
	//Gamename []*GameName `orm:"reverse(many)"`
}



func (u *User) TableName() string{
	return TableName("user")
}

func UserAdd(u *User) (int64,error){
	return orm.NewOrm().Insert(u)
}

func UsergetByOpenid(openid string) (*User,error){
	u := new(User)
	err:= orm.NewOrm().QueryTable(TableName("user")).Filter("openid",openid).One(u)
	if err != nil{
		return nil,err
	}
	return u,err
}

func UsergetByusername(name string) (*User,error){
	u := new(User)
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("user_name",name).One(u)
	if err != nil{
		return nil,err
	}
	return u,err
}

func (u *User) UserUpdate(fileds ...string) error {
	if _,err := orm.NewOrm().Update(u,fileds...);err != nil {
		return err
	}
	return nil
}

func UserDelete(name string) (int64,error){
	query := orm.NewOrm().QueryTable(TableName("user"))
	num,err := query.Filter("user_name",name).Delete()
	return num,err
}

func SearchDataList(pagesize,pageno int) (users []User) {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("user"))
	//if search !=""{
	//	qs=qs.Filter("UserName",search)
	//}
	var us []User
	cnt, err :=  qs.Limit(pagesize, (pageno-1)*pagesize).All(&us)
	if err == nil {
		fmt.Println("count", cnt)
	}
	return us
}

func GetRecordNum() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("user"))
	var us []User
	num, err :=  qs.All(&us)
	if err == nil {
		return num
	}else{
		return 0
	}
}