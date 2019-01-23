package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	self.TplName = "login.html"
}
