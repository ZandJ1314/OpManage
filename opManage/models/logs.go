package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Logs struct {
	Id int
	SourceIp string
	SourceName string
	IsLegalIp bool
	AccessTime time.Time
}

func (l *Logs) TableName() string{
	return TableName("logs")
}

func LogsAdd(l *Logs) (int64,error){
	return orm.NewOrm().Insert(l)
}

func LogsDelete(logintime time.Time) (int64,error){
	query := orm.NewOrm().QueryTable(TableName("logs"))
	num,err := query.Filter("access_time",logintime).Delete()
	return num,err
}

func SearchLogDataList(pagesize,pageno int) (users []Logs) {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("logs"))
	var us []Logs
	cnt, err :=  qs.Limit(pagesize, (pageno-1)*pagesize).All(&us)
	if err == nil {
		fmt.Println("count", cnt)
	}
	return us
}

func GetLogRecordNum() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("logs"))
	var us []Logs
	num, err :=  qs.All(&us)
	if err == nil {
		return num
	}else{
		return 0
	}
}
