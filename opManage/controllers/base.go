package controllers

import (
	"github.com/astaxie/beego"
	"opManage/models"
	"strings"
)


const (
	Msg_OK = 0
	Msg_Err = -1
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName string
	user models.User
	pageSize int
	allowUrl string
}

func (self *BaseController) Prepare(){
	self.pageSize = 20
	controllerName,actionName := self.GetControllerAndAction()//该函数得到两个字串
	//如果是get请求，则是LoginController GET；如果是post请求则是PostControloler POST
	self.controllerName = strings.ToLower(controllerName[0:len(controllerName)-10])//切片
	self.actionName = strings.ToLower(actionName)//字符串转换小写
	self.Data["version"] = beego.AppConfig.String("version") //模板数据，从模板中可以获取
	self.Data["curRoute"] = self.controllerName + "." + self.actionName
	self.Data["curController"] = self.controllerName
	//fmt.Println(self.controllerName,self.actionName)
}



//判断是否为post提交
func (self *BaseController) isPost() bool{
	return self.Ctx.Request.Method == "post"
}


//获取用户IP地址
func (self *BaseController) getClientIp() string{
	s := self.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
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


//加载模板
func (self *BaseController) display(tpl ...string){
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0],"html"},".")
	}else{
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.TplName = tplname
}