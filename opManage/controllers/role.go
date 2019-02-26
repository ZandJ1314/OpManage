package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"html/template"
	"opManage/lib"
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

func (r *RoleController) Get(){
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
		lib.NewLog().Error("role信息提取错误",err)
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
		lib.NewLog().Error("role信息提取错误",err2)
	}
	r.Data["xsrfdata"]=template.HTML(r.XSRFFormHTML())
	r.TplName = "permission/role.html"
}



func (r *RoleController) Post() {
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
		lib.NewLog().Error("role信息提取错误",err)
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
		lib.NewLog().Error("role信息提取错误",err2)
	}
	r.Data["xsrfdata"]=template.HTML(r.XSRFFormHTML())
	r.TplName = "permission/role.html"
}


func (r *RoleController) Prepare(){
	r.EnableXSRF = false
}


func (r *RoleController) ZtreeInfo(){
	sql := "select role_name,roleid,role_pid from op_role;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql).ValuesList(&list)
	if err == nil && res > 0{
		slice := make([]interface{},0)
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["name"] = list[i][0]
			plattest["id"] = list[i][1]
			plattest["pId"] = list[i][2]
			slice = append(slice,plattest)
		}
		r.Data["json"] = slice
		r.ServeJSON()
		return
	}else{
		lib.NewLog().Error("failed",err)
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

func (r *RoleController) AddRole(){
	var user User
	data := r.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		lib.NewLog().Error("json.Unmarshal is err:",err.Error())
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	newGiveRole := new(models.GiveRole)
	newGiveRole.UserName = user.UserName
	newGiveRole.GameName = user.GameName
	//a := reflect.TypeOf(user.RoleInfo).String()
	//fmt.Println(a)  //map[string]interface{}类型
	str,err := json.Marshal(user.RoleInfo)
	roleinfo := string(str)
	newGiveRole.Role = roleinfo
	if _,err := models.GiveRoleAdd(newGiveRole);err != nil{
		lib.NewLog().Error("failed",err)
		r.ajaxMsg(err.Error(),Msg_Err)
	}
	r.ajaxMsg("权限分配成功",Msg_OK)
}

func (r *RoleController) UpdateRole() {
	var users User
	data := r.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &users)
	if err != nil {
		lib.NewLog().Error("json.Unmarshal is err:",err.Error())
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	name := users.UserName
	GiveRole,_ := models.GiveRoleGetByUsername(name)
	GiveRole.GameName = users.GameName
	str,err := json.Marshal(users.RoleInfo)
	roleinfo := string(str)
	GiveRole.Role = roleinfo
	if err := GiveRole.GiveRoleUpdate();err != nil{
		lib.NewLog().Error("failed",err)
		r.ajaxMsg(err.Error(),Msg_Err)
	}else{
		r.ajaxMsg("角色重新分配修改成功",Msg_OK)
	}

}
