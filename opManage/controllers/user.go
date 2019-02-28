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

func (u *UserController) AllUserInfo(){
	name := u.GetString("name")
	head := u.GetString("head")
	u.Data["name"] = name
	u.Data["head"] = head
	//得到当前分页html的数据
	pa,_ := u.GetInt("page")
	pre_page := 10
	//数据总数量
	totals := models.GetRecordNum()
	res := lib.Paginator(pa,pre_page,totals)
	res["name"] = name
	res["head"] = head
	//得到分页的数据
	//fmt.Println(res,pre_page,totals)
	userlist := models.SearchDataList(10,pa)
	//fmt.Println(userlist)
	u.Data["datas"] = userlist
	u.Data["paginator"]=res //分页的数据
	u.Data["totals"] = totals
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
	alaisusername := u.GetString("alaisusername")
	userlevel := u.GetString("userlevel")
	isaddtype := u.GetString("bedStatus")
	var usertype string
	var issuperadministrator int
	if isaddtype == "allot"{
		usertype = u.GetString("addusertype")
		alaisusertype := u.GetString("alaisaddusertype")
		newType := new(models.UserType)
		newType.Usertypename = usertype
		newType.AlaisUsertypename = alaisusertype
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
		newUser.AlaisUserName = alaisusername
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
	GiveRole,_ := models.GiveRoleGetByUsername(username)
	if GiveRole == nil{
		num,err := models.UserDelete(username)
		if num >0 && err == nil{
			msg := username + "已经成功删除了！"
			u.ajaxMsg(msg,Msg_OK)
		}
	}else{
		_,err1 := models.GiveRoleDeleteByUsername(username)
		_,err2 := models.UserDelete(username)
		if err1 == nil && err2 == nil{
			msg := username + "已经成功删除了！"
			u.ajaxMsg(msg,Msg_OK)
		}
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