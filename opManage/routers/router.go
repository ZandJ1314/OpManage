package routers

import (
	"github.com/astaxie/beego"
	"opManage/controllers"
)

func init() {
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/",&controllers.IndexController{})
    beego.Router("/index/info",&controllers.IndexController{},"GET:UserInfo")
    beego.Router("/user",&controllers.UserController{})
    beego.Router("/user/add",&controllers.UserController{},"POST:AddUser")
    beego.Router("/user/delete",&controllers.UserController{},"POST:UserDelete")
    beego.Router("/user/detail",&controllers.UserController{},"POST:Detail")
    beego.Router("/user/judge",&controllers.UserController{},"GET:Judge")
    beego.Router("/user/judgeuser",&controllers.UserController{},"GET:JudgeUser")
    beego.Router("/role",&controllers.RoleController{})
    beego.Router("/role/ztree",&controllers.RoleController{},"POST:ZtreeInfo")
    beego.Router("/role/add",&controllers.RoleController{},"POST:AddRole")
    beego.Router("/role/update",&controllers.RoleController{},"POST:UpdateRole")
    beego.Router("/userquery",&controllers.RoleController{},"GET:UserQuery")
    beego.Router("/permission",&controllers.PermissionController{},"*:PermissionInfo")
    beego.Router("/permission/ztree",&controllers.PermissionController{},"POST:ZtreeInfo")
    beego.Router("/permission/add",&controllers.PermissionController{},"POST:Add")
    beego.Router("/permission/update",&controllers.PermissionController{},"POST:Update")
    beego.Router("/permission/delete",&controllers.PermissionController{},"POST:Delete")
    beego.Router("/permission/selecth",&controllers.PermissionController{},"GET:SelectH")
    beego.Router("/game",&controllers.GameController{},"*:GameInfo")
    beego.Router("/game/add",&controllers.GameController{},"POST:AddGame")
    beego.Router("/game/delete",&controllers.GameController{},"POST:GameDelete")
    beego.Router("/game/detail",&controllers.GameController{},"POST:GameDetail")
    beego.Router("/game/judgegame",&controllers.GameController{},"GET:JudgeGame")
    beego.Router("/game/judgetype",&controllers.GameController{},"GET:JudgeType")
    beego.Router("/system",&controllers.SystemController{})
}

