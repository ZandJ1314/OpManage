package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type GameName struct {
	Id int
	Gamename string
	AlaisGamename string
	GameUrl string
	GamePartment string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);null"`
	//User *User `orm:"rel(fk)"`
	//Gametype *GameType `orm:"rel(fk)"`
	//Role []*Role `orm:"reverse(many)"`
}


func (g *GameName) TableName() string{
	return TableName("gamename")
}

func GameNameAdd(g *GameName) (int64,error) {
	return orm.NewOrm().Insert(g)
}


func GameNameGetByGameid(gameid int) (*GameName, error) {
	g := new(GameName)
	err := orm.NewOrm().QueryTable(TableName("gamename")).Filter("game_id", gameid).One(g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func GameNameGetByGamename(gamename string) (*GameName,error){
	g := new(GameName)
	err := orm.NewOrm().QueryTable(TableName("gamename")).Filter("gamename",gamename).One(g)
	if err != nil{
		return nil,err
	}
	return g,nil
}

func (g *GameName) GamenameUpdate(fields ...string) error {
	if _, err := orm.NewOrm().Update(g, fields...); err != nil {
		return err
	}
	return nil
}

func GameDelete(name string) (int64,error){
	query := orm.NewOrm().QueryTable(TableName("gamename"))
	num,err := query.Filter("gamename",name).Delete()
	return num,err
}

func SearchGameDataList(pagesize,pageno int) (users []GameName) {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("gamename"))
	var us []GameName
	cnt, err :=  qs.Limit(pagesize, (pageno-1)*pagesize).All(&us)
	if err == nil {
		fmt.Println("count", cnt)
	}
	return us
}

func GetGameRecordNum() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("gamename"))
	var us []GameName
	num, err :=  qs.All(&us)
	if err == nil {
		return num
	}else{
		return 0
	}
}