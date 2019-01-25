package controllers

import "github.com/astaxie/beego"

type RegisterController struct {
	beego.Controller
}

func (self *RegisterController) Get() {
	self.TplName = "register/register.html"
}
