package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
)

func Init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	dbhost := beego.AppConfig.String("db.host")
	dbuser := beego.AppConfig.String("db.user")
	dbport := beego.AppConfig.String("db.port")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	dbtimezome := beego.AppConfig.String("db.timezone")
	if dbport == ""{
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Local"
	if dbtimezome == ""{
		dsn = dsn + "&loc=" + url.QueryEscape(dbtimezome)
	}
	orm.RegisterDataBase("default","mysql",dsn)
	orm.RegisterModel(new(User),new(Role),new(GameType),new(GameName),new(GiveRole),new(UserType),new(Logs))
	orm.RunSyncdb("default",false,true)
	//orm.RegisterDataBase("db2","mysql",dsn2)
	fmt.Println("连接数据库成功")
	if beego.AppConfig.String("runmode") == "dev"{
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
