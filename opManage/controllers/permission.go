package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"html/template"
	"opManage/libs"
	"opManage/models"
	"strconv"
	"strings"
)

type PermissionController struct {
	BaseController
}

func (p *PermissionController) PermissionInfo(){
	name := p.GetString("name")
	head := p.GetString("head")
	p.Data["name"] = name
	p.Data["head"] = head
	sql := "select role from op_role;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql).ValuesList(&list)
	slice := make([]interface{},0)
	if err == nil && res > 0{
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["rolename"] = list[i][0]
			slice = append(slice,plattest)
		}
		p.Data["result"] = slice
	}
	p.Data["xsrfdata"]=template.HTML(p.XSRFFormHTML())
	p.TplName = "permission/permission.html"
}

func (p *PermissionController) Prepare(){
	p.EnableXSRF = false
}


func (p *PermissionController) ZtreeInfo(){
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
		p.Data["json"] = slice
		p.ServeJSON()
		return
	}else{
		libs.NewLog().Error("failed",err)
	}
}


func (p *PermissionController) Add(){
	role := p.GetString("baserole")
	rolename := p.GetString("role_name")
	higherrole := p.GetString("heigher_role")
	//rolelevel := p.GetString("role_level")
	//gamename := p.GetString("gamename")
	//GameName,_ := models.GameNameGetByGamename(gamename)
	if role == "所有权限"{
		roleid := 0
		rolepid := 0
		newRole := new(models.Role)
		newRole.Role = role
		newRole.RoleName = rolename
		newRole.RoleLevel = "0"
		newRole.HigherRole = higherrole
		newRole.Roleid = roleid
		newRole.RolePid = rolepid
		//newRole.GameName = GameName
		if _,err := models.RoleAdd(newRole);err != nil{
			libs.NewLog().Error("failed",err)
			p.ajaxMsg(err.Error(),Msg_Err)
		}
		p.ajaxMsg("权限增加成功",Msg_OK)

	}else{
		Role,err := models.RoleGetByHigherRole(higherrole)
		rolelevel := Role.RoleLevel
		nrolelevel,_ := strconv.Atoi(rolelevel)
		nrolelevel = nrolelevel + 1
		str_rolelevel := strconv.Itoa(nrolelevel)
		newrolelevel,_ := strconv.ParseUint(str_rolelevel,0,64)
		//rolelevel = uint64(rolelevel)
		//rolelevel,_ := strconv.Atoi(rolelevel)
		var newroleid uint64
		libs.NewLog().Error("failed",err)
		higherroleid := Role.Roleid
		rolepid := higherroleid
		sql := "select roleid from op_role where role_pid = ?"
		o := orm.NewOrm()
		var list []orm.ParamsList
		res,err := o.Raw(sql,higherroleid).ValuesList(&list)
		if err == nil && res > 0 {
			var roleid interface{}
			for i:=0;i<len(list);i++{
				roleid = list[i][0]
			}
			oldroleid,_:= roleid.(string)
			roleidold,_ := strconv.Atoi(oldroleid)
			roleidnew := roleidold + 1
			newroleid = uint64(roleidnew)

		}else{
			level := newrolelevel-1
			newroleid = libs.Exponet(10,level) //数据类型是unit64
			newroleid = newroleid * uint64(rolepid)
			libs.NewLog().Error("failed",err)
		}
		newRole := new(models.Role)
		newRole.Role = role
		newRole.RoleName = rolename
		newRole.RoleLevel = str_rolelevel
		newRole.HigherRole = higherrole
		newRole.Roleid = int(newroleid)
		newRole.RolePid = rolepid
		//newRole.GameName = GameName
		if _,err := models.RoleAdd(newRole);err != nil{
			libs.NewLog().Error("failed",err)
			p.ajaxMsg(err.Error(),Msg_Err)
		}
		p.ajaxMsg("权限增加成功",Msg_OK)

	}
}

func (p *PermissionController) Update()  {
	oldrolename := p.GetString("oldrolename")
	Role,_ := models.RoleGetByroleName(oldrolename)
	//oldhigherrole := Role.HigherRole
	//olddetail := Role.RoleDescription
	higherrole := p.GetString("higherrole")
	newrolename := p.GetString("newrolename")
	basenewrolename := p.GetString("basenewrolename")
	if newrolename == "" || basenewrolename == ""{
		result := make(map[string]interface{})
		result["message"] = "输入框不能为空！！"
		p.Data["json"] = result
		p.ServeJSON()
	}else{
		if higherrole == "所有权限"{
			Role.Role = newrolename
			Role.RoleName = basenewrolename
			if err := Role.RoleUpdate();err != nil{
				libs.NewLog().Error("failed",err)
				p.ajaxMsg(err.Error(),Msg_Err)
			}else{
				p.ajaxMsg("权限修改成功",Msg_OK)
			}
		}else{
			HigherRole,_ := models.RoleGetByroleName(higherrole)
			rolelevel := HigherRole.RoleLevel
			nrolelevel,_ := strconv.Atoi(rolelevel)
			nrolelevel = nrolelevel + 1
			str_rolelevel := strconv.Itoa(nrolelevel)
			newrolelevel,_ := strconv.ParseUint(str_rolelevel,0,64)
			var newroleid uint64
			higherroleid := HigherRole.Roleid
			rolepid := higherroleid
			sql := "select roleid from op_role where role_pid = ?"
			o := orm.NewOrm()
			var list []orm.ParamsList
			res,err := o.Raw(sql,higherroleid).ValuesList(&list)
			if err == nil && res > 0 {
				var roleid interface{}
				for i:=0;i<len(list);i++{
					roleid = list[i][0]
				}
				oldroleid,_:= roleid.(string)
				roleidold,_ := strconv.Atoi(oldroleid)
				roleidnew := roleidold + 1
				newroleid = uint64(roleidnew)

			}else{
				level := newrolelevel-1
				newroleid = libs.Exponet(10,level) //数据类型是unit64
				newroleid = newroleid * uint64(rolepid)
				libs.NewLog().Error("failed",err)
			}
			Role.Role = newrolename
			Role.RoleName = basenewrolename
			Role.RoleLevel = str_rolelevel
			Role.HigherRole = higherrole
			Role.Roleid = int(newroleid)
			Role.RolePid = rolepid
			if err := Role.RoleUpdate();err != nil{
				libs.NewLog().Error("failed",err)
				p.ajaxMsg(err.Error(),Msg_Err)
			}else{
				p.ajaxMsg("权限修改成功",Msg_OK)
			}
		}
	}

}

func (p *PermissionController) Delete()  {
	rolename := p.GetString("rolename")
	Role,_ := models.RoleGetByroleName(rolename)
	rolepid := Role.Roleid
	sql := "select roleid from op_role where role_pid = ?"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql,rolepid).ValuesList(&list)
	if err == nil && res > 0 {
		var number int
		for i:=0;i<len(list);i++{
			roleid := list[i][0]
			newroleid,_ := roleid.(int)
			num,err := Role.RoleDeleteByRoleid(newroleid)
			if err != nil && num <= 0{
				libs.NewLog().Error("failed",err)
				p.ajaxMsg(err.Error(),Msg_Err)
			}else{
				number ++
			}
		}
		msg := "权限删除成功，共删除" + string(number) + "条"
		p.ajaxMsg(msg,Msg_OK)
	}else{
		num,err := Role.RoleDeleteByRolename(rolename)
		if err != nil && num <= 0{
			libs.NewLog().Error("failed",err)
			p.ajaxMsg(err.Error(),Msg_Err)
		}else{
			p.ajaxMsg("权限删除成功,共删除一条",Msg_OK)
		}
	}

}

func (p *PermissionController) SelectH(){
	rolename := p.GetString("name")
	newrolename := strings.Split(rolename,"(")[0]
	Role,_ := models.RoleGetByroleName(newrolename)
	heighername := Role.HigherRole
	role_name := Role.RoleName
	plattest := make(map[string]interface{})
	plattest["name"] = heighername
	plattest["rolename"] = role_name
	p.Data["json"] = &plattest
	p.ServeJSON()
}
