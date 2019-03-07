package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"html/template"
	"opManage/libs"
	"opManage/models"
)

type RoleController struct {
	BaseController
}

type User struct {
	UserName string
	GameName string
	RoleInfo interface{}
}

type UserType struct {
	Name string
}

func (r *RoleController) AllRoleInfo(){
	name := r.GetString("name")
	head := r.GetString("head")
	r.Data["name"] = name
	r.Data["head"] = head
	sql := "select usertypename from op_usertype;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql).ValuesList(&list)
	if err == nil && res > 0{
		slice := make([]interface{},0)
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["usertypename"] = list[i][0]
			slice = append(slice,plattest)
		}
		r.Data["result"] = slice
	}else {
		libs.NewLog().Error("usertype信息提取错误",err)
	}
	sql2 := "select gamename from op_gamename;"
	o2 := orm.NewOrm()
	var list2 []orm.ParamsList
	res2,err2 := o2.Raw(sql2).ValuesList(&list2)
	if err2 == nil && res2 > 0{
		slice2 := make([]interface{},0)
		for i := 0;i<len(list2);i++{
			plattest2 := make(map[string]interface{})
			plattest2["gamename"] = list2[i][0]
			slice2 = append(slice2,plattest2)
		}
		r.Data["result2"] = slice2
	}else {
		libs.NewLog().Error("game信息提取错误",err2)
	}
	sql3 := "select user_name from op_user;"
	o3 := orm.NewOrm()
	var list3 []orm.ParamsList
	res3,err3 := o3.Raw(sql3).ValuesList(&list3)
	if err3 == nil && res3 > 0 {
		slice3 := make([]interface{},0)
		for i := 0;i<len(list3);i++{
			plattest3 := make(map[string]interface{})
			plattest3["username"] = list3[i][0]
			slice3 = append(slice3,plattest3)
		}
		r.Data["result3"] = slice3
	}else{
		libs.NewLog().Error("user信息提取错误",err2)
	}
	r.Data["xsrfdata"]=template.HTML(r.XSRFFormHTML())
	r.TplName = "permission/role.html"
}



func (r *RoleController) Prepare(){
	r.EnableXSRF = false
}


func (r *RoleController) ZtreeInfo(){
	sql := "select role,role_name,roleid,role_pid from op_role;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql).ValuesList(&list)
	if err == nil && res > 0{
		slice := make([]interface{},0)
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			name,_ := list[i][0].(string)
			rolename,_ := list[i][1].(string)
			plattest["name"] = name + "("+rolename+")"
			plattest["id"] = list[i][2]
			plattest["pId"] = list[i][3]
			slice = append(slice,plattest)
		}
		r.Data["json"] = slice
		r.ServeJSON()
		return
	}else{
		libs.NewLog().Error("failed",err)
	}
}

func (r *RoleController) UserQuery() {
	r.TplName = "permission/role.html"
	name := r.GetString("typename")
	sql := "select user_name from op_user where department = ?;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql,name).ValuesList(&list)
	slice := make([]interface{},0)
	//var request string
	if err == nil && res > 0{
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["name"] = list[i][0]
			slice = append(slice,plattest)
		}
		r.Data["json"] = &slice
		r.ServeJSON()
	}else{
		plattest := make(map[string]interface{})
		plattest["name"] = ""
		r.Data["json"] = &plattest
		r.ServeJSON()
	}

}
func (r *RoleController) GameQuery() {
	r.TplName = "permission/role.html"
	name := r.GetString("username")
	sql := "select game_name from op_giverole where user_name = ?;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql,name).ValuesList(&list)
	slice := make([]interface{},0)
	//var request string
	if err == nil && res > 0{
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["name"] = list[i][0]
			slice = append(slice,plattest)
		}
		r.Data["json"] = &slice
		r.ServeJSON()
	}else{
		plattest := make(map[string]interface{})
		plattest["name"] = ""
		r.Data["json"] = &plattest
		r.ServeJSON()
	}

}


func (r *RoleController) AddRole(){
	var user User
	data := r.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		libs.NewLog().Error("json.Unmarshal is err:",err.Error())
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	gamename := user.GameName
	username := user.UserName
	sql := "select role from op_giverole where user_name = ? and game_name = ?"
	selectIds := []string{username,gamename}
	o := orm.NewOrm()
	var maps []orm.Params
	res,err1 := o.Raw(sql,selectIds).Values(&maps)
	if err1 == nil && res > 0{
		result := make(map[string]interface{})
		result["message"] = "您已经给"+username+"分配过"+gamename+"的权限了，请选择其他游戏"
		r.Data["json"] = result
		r.ServeJSON()
	}else {
		newGiveRole := new(models.GiveRole)
		newGiveRole.UserName = username
		newGiveRole.GameName = gamename
		//a := reflect.TypeOf(user.RoleInfo).String()
		//fmt.Println(a)  //map[string]interface{}类型
		str,_ := json.Marshal(user.RoleInfo)
		roleinfo := string(str)
		if roleinfo == "{}"{
			result := make(map[string]interface{})
			result["message"] = "您没有选择绑定任何权限，请您选择之后再提交!!"
			r.Data["json"] = result
			r.ServeJSON()
		}else{
			newGiveRole.Role = roleinfo
			if _,err := models.GiveRoleAdd(newGiveRole);err != nil{
				libs.NewLog().Error("failed",err)
				r.ajaxMsg(err.Error(),Msg_Err)
			}
			r.ajaxMsg("权限分配成功",Msg_OK)
		}
	}
}

func (r *RoleController) UpdateRole() {
	var users User
	data := r.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &users)
	if err != nil {
		libs.NewLog().Error("json.Unmarshal is err:",err.Error())
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	name := users.UserName
	GiveRole,_ := models.GiveRoleGetByUsername(name)
	GiveRole.GameName = users.GameName
	str,err := json.Marshal(users.RoleInfo)
	roleinfo := string(str)
	if roleinfo == "{}"{
		result := make(map[string]interface{})
		result["message"] = "您没有选择绑定任何权限，请您选择之后再提交!!"
		r.Data["json"] = result
		r.ServeJSON()
	}else{
		GiveRole.Role = roleinfo
		if err := GiveRole.GiveRoleUpdate();err != nil{
			libs.NewLog().Error("failed",err)
			r.ajaxMsg(err.Error(),Msg_Err)
		}else{
			r.ajaxMsg("角色重新分配修改成功",Msg_OK)
		}
	}
}

func (r *RoleController) DeleteRole() {
	username := r.GetString("deletename")
	gamename := r.GetString("gamerole")
	fmt.Println(username,gamename)
	num,err := models.GiveRoleDeleteByGameAndUser(username,gamename)
	fmt.Println(num,err)
	if err == nil && num > 0{
		r.ajaxMsg("角色删除成功",Msg_OK)
	}else{
		r.ajaxMsg("角色删除失败",Msg_Err)
	}
}
