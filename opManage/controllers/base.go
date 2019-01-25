package controllers

import "github.com/astaxie/beego"


const (
	Msg_OK = 0
	Msg_Err = -1
)

type BaseController struct {
	beego.Controller
}


//定义重定向
func (self *BaseController) redirect(url string){
	self.Redirect(url,302)
	self.StopRun()
}

//ajax返回数据
func (self *BaseController) ajaxMsg(msg interface{},msgno int){
	result := make(map[string]interface{})
	result["status"] = msgno
	result["message"] = msg
	self.Data["json"] = result
	self.ServeJSON()
	//self.StopRun()
}
