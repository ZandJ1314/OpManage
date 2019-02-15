package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type GameName struct {
	Id int
	GameId int
	Gamename string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);null"`
	User *User `orm:"rel(fk)"`
	Gametype *GameType `orm:"rel(fk)"`
	Role []*Role `orm:"reverse(many)"`
}


func (g *GameName) TableName() string{
	return TableName("gamename")
}

func GameNameAdd(g *GameType) (int64,error) {
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