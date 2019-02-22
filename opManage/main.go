package main

import (
	"opManage/models"
	_ "opManage/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

