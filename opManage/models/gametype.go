package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type GameType struct {
	Id int
	Typepid int
	Gametypename string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);null"`
	User *User `orm:"rel(fk)"`  //设置一对多的关系
	Gamename []*GameName `orm:"reverse(many)"`
}


func (g *GameType) TableName() string{
	return TableName("gametype")
}

func GameTypeAdd(g *GameType) (int64,error) {
	return orm.NewOrm().Insert(g)
}


func GameTypeGetById(typepid int) (*GameType, error) {
	g := new(GameType)
	err := orm.NewOrm().QueryTable(TableName("gametype")).Filter("typepid", typepid).One(g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *GameType) GametypeUpdate(fields ...string) error {
	if _, err := orm.NewOrm().Update(g, fields...); err != nil {
		return err
	}
	return nil
}