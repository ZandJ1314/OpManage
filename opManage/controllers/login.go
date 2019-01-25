package controllers

type LoginController struct {
	BaseController
}

func (self *LoginController) Get() {
	self.TplName = "login/login.html"
}
