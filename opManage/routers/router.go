package routers

import (
	"github.com/astaxie/beego"
	"opManage/controllers"
)

func init() {
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/",&controllers.IndexController{})
    beego.Router("/user",&controllers.UserController{})
    beego.Router("/role",&controllers.RoleController{})
    beego.Router("/permission",&controllers.PermissionController{})
    beego.Router("/permission/ztree",&controllers.PermissionController{},"POST:ZtreeInfo")
    beego.Router("/permission/add",&controllers.PermissionController{},"POST:Add")
    beego.Router("/permission/update",&controllers.PermissionController{},"POST:Update")
    beego.Router("/permission/delete",&controllers.PermissionController{},"POST:Delete")
    beego.Router("/game",&controllers.GameController{})
    beego.Router("/system",&controllers.SystemController{})
}
