package routers

import (
	"github.com/astaxie/beego"
	"opManage/controllers"
)

func init() {
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/",&controllers.IndexController{})
}
