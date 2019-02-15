package controllers

import (
	"github.com/astaxie/beego/orm"
	"opManage/lib"
	"opManage/models"
	"strconv"
)

type PermissionController struct {
	BaseController
}

func (p *PermissionController) Get(){
	p.TplName = "permission/permission.html"
}

func (p *PermissionController) Post()  {
	p.TplName = "permission/permission.html"
}


func (p *PermissionController) ZtreeInfo(){
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
		p.Data["json"] = slice
		p.ServeJSON()
		return
	}else{
		lib.NewLog().Error("failed",err)
	}
}


func (p *PermissionController) Add(){
	rolename := p.GetString("role_name")
	roledescription := p.GetString("detail")
	higherrole := p.GetString("heigher_role")
	rolelevel := p.GetString("role_level")
	gamename := p.GetString("gamename")
	GameName,_ := models.GameNameGetByGamename(gamename)
	if rolelevel == "0"{
		roleid := 0
		rolepid := 0
		newRole := new(models.Role)
		newRole.RoleName = rolename
		newRole.RoleDescription = roledescription
		newRole.RoleLevel = rolelevel
		newRole.HigherRole = higherrole
		newRole.Roleid = roleid
		newRole.RolePid = rolepid
		newRole.GameName = GameName
		if _,err := models.RoleAdd(newRole);err != nil{
			lib.NewLog().Error("failed",err)
			p.ajaxMsg(err.Error(),Msg_Err)
		}
		p.ajaxMsg("权限增加成功",Msg_OK)

	}else{
		rolelevel,_ := strconv.ParseUint(rolelevel,0,64)
		//rolelevel,_ := strconv.Atoi(rolelevel)
		var newroleid uint64
		Role,err := models.RoleGetByHigherRole(higherrole)
		lib.NewLog().Error("failed",err)
		higherroleid := Role.Roleid
		rolepid := higherroleid
		sql := "select roleid from op_role where role_pid = ?"
		o := orm.NewOrm()
		var list []orm.ParamsList
		res,err := o.Raw(sql,higherroleid).ValuesList(&list)
		if err == nil && res > 0 {
			var roleid interface{}
			if len(list) == 1{
				level := rolelevel-1
				newroleid = lib.Exponet(10,level)//数据类型是unit64
				//a := reflect.TypeOf(newroleid).String()//判断数据类型
				//fmt.Println(a)
			}else{
				for i:=0;i<len(list);i++{
					roleid = list[i][0]
				}
				oldroleid,_:= roleid.(string)
				roleidold,_ := strconv.Atoi(oldroleid)
				roleidnew := roleidold + 1
				newroleid = uint64(roleidnew)
			}
			newRole := new(models.Role)
			newRole.RoleName = rolename
			newRole.RoleDescription = roledescription
			newRole.RoleLevel = string(rolelevel)
			newRole.HigherRole = higherrole
			newRole.Roleid = int(newroleid)
			newRole.RolePid = rolepid
			newRole.GameName = GameName
			if _,err := models.RoleAdd(newRole);err != nil{
				lib.NewLog().Error("failed",err)
				p.ajaxMsg(err.Error(),Msg_Err)
			}
			p.ajaxMsg("权限增加成功",Msg_OK)

		}else{
			lib.NewLog().Error("failed",err)
		}

	}
}

func (p *PermissionController) Update()  {
	oldrolename := p.GetString("oldrolename")
	Role,_ := models.RoleGetByroleName(oldrolename)
	//oldhigherrole := Role.HigherRole
	//olddetail := Role.RoleDescription
	higherrole := p.GetString("higherrole")
	newrolename := p.GetString("newrolename")
	newdetail := p.GetString("newdetail")
	if higherrole == ""{
		Role.RoleName = newrolename
		Role.RoleDescription = newdetail
		if err := Role.RoleUpdate();err != nil{
			lib.NewLog().Error("failed",err)
			p.ajaxMsg(err.Error(),Msg_Err)
		}else{
			p.ajaxMsg("权限修改成功",Msg_OK)
		}
	}else{
		Role.RoleName = newrolename
		Role.RoleDescription = newdetail
		HigherRole,_ := models.RoleGetByroleName(higherrole)
		rolelevel := HigherRole.RoleLevel
		newlevel,_ := strconv.ParseUint(rolelevel,0,64)
		//newlevel,_ := strconv.Atoi(rolelevel)
		var newroleid uint64
		newlevel = newlevel - 1
		Role.RoleLevel = string(newlevel)
		Role.RolePid = HigherRole.Roleid
		sql := "select roleid from op_role where role_pid = ?"
		o := orm.NewOrm()
		var list []orm.ParamsList
		res,err := o.Raw(sql,HigherRole.Roleid).ValuesList(&list)
		if err == nil && res > 0 {
			var roleid interface{}
			if len(list) == 1{
				level := newlevel-1
				newroleid = lib.Exponet(10,level)//数据类型是unit64
				//a := reflect.TypeOf(newroleid).String()//判断数据类型
				//fmt.Println(a)
			}else{
				for i:=0;i<len(list);i++{
					roleid = list[i][0]
				}
				oldroleid,_:= roleid.(string)
				roleidold,_ := strconv.Atoi(oldroleid)
				roleidnew := roleidold + 1
				newroleid = uint64(roleidnew)
			}
		}else{
			lib.NewLog().Error("failed",err)
		}
		Role.Roleid = int(newroleid)
		if err := Role.RoleUpdate();err != nil{
			lib.NewLog().Error("failed",err)
			p.ajaxMsg(err.Error(),Msg_Err)
		}else{
			p.ajaxMsg("权限修改成功",Msg_OK)
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
				lib.NewLog().Error("failed",err)
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
			lib.NewLog().Error("failed",err)
			p.ajaxMsg(err.Error(),Msg_Err)
		}else{
			p.ajaxMsg("权限删除成功,共删除一条",Msg_OK)
		}
	}

}