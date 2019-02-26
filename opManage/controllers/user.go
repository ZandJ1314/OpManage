package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"html/template"
	"opManage/lib"
	"opManage/models"
)

type UserController struct {
	BaseController
}

type DelUser struct {
	Name string
}

func (u *UserController) Get(){
	name := u.GetString("name")
	head := u.GetString("head")
	u.Data["name"] = name
	u.Data["head"] = head
	sql := "select openid,user_name,department,phone_number,email_adress,head_portrait_name,issuperadministrator,create_time from op_user;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql).ValuesList(&list)
	if err == nil && res > 0{
		slice := make([]interface{},0);
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["openid"] = list[i][0]
			plattest["username"] = list[i][1]
			plattest["department"] = list[i][2]
			plattest["phonenumber"] = list[i][3]
			plattest["emailadress"] = list[i][4]
			plattest["headportraitname"] = list[i][5]
			plattest["issuperadministrator"] = list[i][6]
			plattest["createtime"] = list[i][7]
			//plattest["index"] = "t"+ strconv.Itoa(i)
			slice = append(slice,plattest)
		}
		u.Data["result"] = slice

	}else{
		lib.NewLog().Error("user信息提取错误",err)
	}
	sql2 := "select usertypename from op_usertype;"
	o2 := orm.NewOrm()
	var list2 []orm.ParamsList
	res2,err2 := o2.Raw(sql2).ValuesList(&list2)
	if err2 == nil && res2 > 0 {
		slice2 := make([]interface{},0)
		for i := 0;i<len(list2);i++{
			plattest2 := make(map[string]interface{})
			plattest2["usertypename"] = list2[i][0]
			slice2 = append(slice2,plattest2)
		}
		u.Data["result2"] = slice2
	}else{
		lib.NewLog().Error("usertype信息提取错误",err2)
	}
	u.Data["xsrfdata"]=template.HTML(u.XSRFFormHTML())
	u.TplName = "permission/user.html"
}

func (u *UserController) Post(){
	name := u.GetString("name")
	head := u.GetString("head")
	u.Data["name"] = name
	u.Data["head"] = head
	sql := "select openid,user_name,department,phone_number,email_adress,head_portrait_name,issuperadministrator,create_time from op_user;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql).ValuesList(&list)
	if err == nil && res > 0{
		slice := make([]interface{},0);
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["openid"] = list[i][0]
			plattest["username"] = list[i][1]
			plattest["department"] = list[i][2]
			plattest["phonenumber"] = list[i][3]
			plattest["emailadress"] = list[i][4]
			plattest["headportraitname"] = list[i][5]
			plattest["issuperadministrator"] = list[i][6]
			plattest["createtime"] = list[i][7]
			slice = append(slice,plattest)
		}
		u.Data["result"] = slice

	}else{
		lib.NewLog().Error("user信息提取错误",err)
	}
	sql2 := "select usertypename from op_usertype;"
	o2 := orm.NewOrm()
	var list2 []orm.ParamsList
	res2,err2 := o2.Raw(sql2).ValuesList(&list2)
	if err2 == nil && res2 > 0 {
		slice2 := make([]interface{},0)
		for i := 0;i<len(list2);i++{
			plattest2 := make(map[string]interface{})
			plattest2["usertypename"] = list2[i][0]
			slice2 = append(slice2,plattest2)
		}
		u.Data["result2"] = slice2
	}else{
		lib.NewLog().Error("usertype信息提取错误",err2)
	}
	u.Data["xsrfdata"]=template.HTML(u.XSRFFormHTML())
	u.TplName = "permission/user.html"
}


func (u *UserController) AddUser(){
	username := u.GetString("username")
	userlevel := u.GetString("userlevel")
	isaddtype := u.GetString("bedStatus")
	var usertype string
	var issuperadministrator int
	if isaddtype == "allot"{
		usertype = u.GetString("addusertype")
		newType := new(models.UserType)
		newType.Usertypename = usertype
		if _,err := models.UserTypeAdd(newType);err != nil{
			lib.NewLog().Error("failed",err)
		}
	}else{
		usertype = u.GetString("usertype")
	}
	if userlevel == "普通管理员"{
		issuperadministrator = 0
	}else{
		issuperadministrator = 1
	}
	if username == "" || usertype == ""{
		result := make(map[string]interface{})
		result["message"] = "输入框不能为空！！"
		u.Data["json"] = result
		u.ServeJSON()
	}else{
		//newUserType,_ := models.UserTypegetByuserTypename(usertype)
		newUser := new(models.User)
		newUser.UserName = username
		newUser.Department = usertype
		newUser.ManageName = userlevel
		newUser.Issuperadministrator = issuperadministrator
		//newUser.UserType = newUserType
		if _,err := models.UserAdd(newUser);err != nil{
			lib.NewLog().Error("failed",err)
			u.ajaxMsg(err.Error(),Msg_Err)
		}
		u.ajaxMsg("管理员添加成功",Msg_OK)
	}


}

func (u *UserController) Prepare(){
	u.EnableXSRF = false
}


func (u *UserController) UserDelete(){
	var deluser DelUser
	data := u.Ctx.Input.RequestBody
	err := json.Unmarshal(data,&deluser)
	if err != nil{
		lib.NewLog().Error("json.Unmarshal is err:",err.Error())
	}
	username := deluser.Name
	num,err := models.UserDelete(username)
	if num >0 && err == nil{
		msg := username + "已经成功删除了！"
		u.ajaxMsg(msg,Msg_OK)
	}
}

func (u *UserController) Detail()  {
	var detuser DelUser
	data := u.Ctx.Input.RequestBody
	err := json.Unmarshal(data,&detuser)
	if err != nil{
		lib.NewLog().Error("json.Unmarshal is err:",err.Error())
	}
	username := detuser.Name
	sql := "select game_name from op_giverole where user_name = ?"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql,username).ValuesList(&list)
	slice := make([]interface{},0)
	if err == nil && res > 0{
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["gamename"] = list[i][0]
			slice = append(slice,plattest)
		}
		u.Data["json"] = slice
	}else{
		plattest := make(map[string]interface{})
		plattest["gamename"] = "该管理员暂无游戏可以管理！"
		u.Data["json"] = plattest
		lib.NewLog().Error("gamename信息提取错误",err)
	}
	u.ServeJSON()
}


func (u *UserController) Judge(){
	name := u.GetString("name")
	_,err := models.UserTypegetByuserTypename(name)
	plattest := make(map[string]interface{})
	if err == nil{
		plattest["name"] = "false"
	}else{
		plattest["name"] = "true"
	}
	u.Data["json"] = &plattest
	u.ServeJSON()
}

func (u *UserController) JudgeUser() {
	name := u.GetString("name")
	_,err := models.UsergetByusername(name)
	plattest := make(map[string]interface{})
	if err == nil{
		plattest["name"] = "false"
	}else{
		plattest["name"] = "true"
	}
	u.Data["json"] = &plattest
	u.ServeJSON()
}