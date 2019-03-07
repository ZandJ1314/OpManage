package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/gomodule/redigo/redis"
	"html/template"
	"opManage/libs"
	"opManage/models"
	"strings"
	"time"
)

type IndexController struct {
	BaseController
}

type Users struct {
	Name string
	HeadPortraitName string
	ManageName string
	Issuperadministrator int
	//GameInfo interface{}
}


func (r *IndexController) Get() {
	code := r.GetString("code")
	access_token := libs.GetAccessToken()
	persistent_code,openid1 := libs.GetPerpetualCode(code,access_token)
	sns_token := libs.GetSnsToken(persistent_code,openid1,access_token)
	var name1,openid string
	name1,openid = libs.GetUserInfo(sns_token)
	logintime := time.Now()
	clientIp := r.getClientIp()
	newlogs := new(models.Logs)
	newlogs.SourceIp = clientIp
	newlogs.SourceName = name1
	newlogs.AccessTime = logintime
	if clientIp == "221.237.152.208"{
		newlogs.IsLegalIp = true
	}else{
		newlogs.IsLegalIp = false
	}
	if _,err := models.LogsAdd(newlogs);err != nil{
		libs.NewLog().Error("failed",err)
	}
	libs.StartTimer(logintime)
	name := r.Ctx.GetCookie("username")
	//sessionid := r.Ctx.GetCookie("sessionid")
	redis,err := cache.NewCache("redis",`{"conn":"127.0.0.1:6379", "key":"cacheuserinfo"}`)
	if name1 == "error"{
		if err != nil{
			libs.NewLog().Error("failed","缓存数据存取错误")
		}
		//sessionid := redis.Get("sessionid")
		//if sessionid == nil{
		//	r.redirect("/login")
		//}
		sessionid := r.Ctx.GetCookie("sessionid")
		userinfo := redis.Get(sessionid)
		fmt.Println("hello")
		fmt.Println(userinfo)
		var users Users
		err := json.Unmarshal(userinfo.([]byte),&users)
		if err != nil{
			libs.NewLog().Error("redis.failed",err)
		}
		newuserinfo := make(map[string]interface{})
		newuserinfo["name"] = users.Name
		newuserinfo["headPortraitName"] = users.HeadPortraitName
		newuserinfo["managename"] = users.ManageName
		newuserinfo["issuperadministrator"] = users.Issuperadministrator
		//newuserinfo["gameinfo"] = users.GameInfo
		r.Data["xsrfdata"]=template.HTML(r.XSRFFormHTML())
		r.Data["result"] = newuserinfo
		r.TplName = "login/index.html"
		//a := reflect.TypeOf(userinfo).String()//判断数据类型
	}else{
		if name == "" || name != name1{
			//if name == "" && sessionid == ""{
			//	r.redirect("/login")
			//}else {
			r.Ctx.SetCookie("username",name1,1800)
			r.Ctx.SetCookie("sessionid",openid,1800)
			userinfo := make(map[string]interface{})
			//if dataurl[lens-5:] == "LOGIN"{
			User,err := models.UsergetByusername(name1)
			if err != nil{
				libs.NewLog().Error("failed",err)
				r.Data["result"] = "您属于非法用户，请与管理员确认您的身份信息"
				r.TplName = "login/login.html"
			}else{
				User.Openid = openid
				if err := User.UserUpdate();err != nil{
					libs.NewLog().Error("failed",err)
				}
				name := User.UserName
				userinfo["name"] = name
				headPortraitName := User.HeadPortraitName
				userinfo["headPortraitName"] =headPortraitName
				manageName := User.ManageName
				userinfo["managename"] = manageName
				issuperadministrator := User.Issuperadministrator
				userinfo["issuperadministrator"] = issuperadministrator
				if err != nil{
					libs.NewLog().Error("failed","缓存数据添加错误")
				}
				//将map类型转换为json字符串，然后存入在redis中，后面直接转换就行
				value,_ := json.Marshal(userinfo)
				redis.Put(openid,value,time.Second*1800)
				r.Data["xsrfdata"]=template.HTML(r.XSRFFormHTML())
				r.Data["result"] = userinfo
				r.TplName = "login/index.html"
			}

		}else{
			//redis,err := cache.NewCache("redis",`{"conn":"127.0.0.1:6379", "key":"cacheuserinfo"}`)
			if err != nil{
				libs.NewLog().Error("failed","缓存数据存取错误")
			}
			//sessionid := redis.Get("sessionid")
			//if sessionid == nil{
			//	r.redirect("/login")
			//}
			userinfo := redis.Get(openid)
			//sessionid = string(sessionid.([]byte))
			//a := reflect.TypeOf(userinfo).String()//判断数据类型
			var users Users
			err := json.Unmarshal(userinfo.([]byte),&users)
			if err != nil{
				libs.NewLog().Error("redis.failed",err)
			}
			newuserinfo := make(map[string]interface{})
			newuserinfo["name"] = users.Name
			newuserinfo["headPortraitName"] = users.HeadPortraitName
			newuserinfo["managename"] = users.ManageName
			newuserinfo["issuperadministrator"] = users.Issuperadministrator
			//newuserinfo["gameinfo"] = users.GameInfo
			r.Data["xsrfdata"]=template.HTML(r.XSRFFormHTML())
			r.Data["result"] = newuserinfo
			r.TplName = "login/index.html"
		}
	}
}

func (r *IndexController) Post(){
	openid := r.Ctx.GetCookie("sessionid")
	redis,err := cache.NewCache("redis",`{"conn":"127.0.0.1:6379", "key":"cacheuserinfo"}`)
	if err != nil{
		libs.NewLog().Error("failed","缓存数据存取错误")
	}
	userinfo := redis.Get(openid)
	//a := reflect.TypeOf(userinfo).String()//判断数据类型
	var users Users
	err1 := json.Unmarshal(userinfo.([]byte),&users)
	if err != nil{
		libs.NewLog().Error("redis.failed",err1)
	}
	newuserinfo := make(map[string]interface{})
	name := users.Name
	sql := "select head_portrait_name from op_user where openid = ?"
	o := orm.NewOrm()
	var maps []orm.Params
	res,err := o.Raw(sql,openid).Values(&maps)
	if err == nil && res > 0{
		newuserinfo["headPortraitName"] = maps[0]["head_portrait_name"]
		userinfo := make(map[string]interface{})
		opUser,_ := models.UsergetByOpenid(openid)
		userinfo["name"] = name
		userinfo["headPortraitName"] =maps[0]["head_portrait_name"]
		userinfo["managename"] = opUser.ManageName
		userinfo["issuperadministrator"] = opUser.Issuperadministrator
		if err != nil{
			libs.NewLog().Error("failed","缓存数据添加错误")
		}
		//将map类型转换为json字符串，然后存入在redis中，后面直接转换就行
		value,_ := json.Marshal(userinfo)
		redis.Put(openid,value,time.Second*1800)
	}
	newuserinfo["name"] = users.Name
	newuserinfo["managename"] = users.ManageName
	newuserinfo["issuperadministrator"] = users.Issuperadministrator
	//newuserinfo["gameinfo"] = users.GameInfo
	r.Data["xsrfdata"]=template.HTML(r.XSRFFormHTML())
	r.Data["result"] = newuserinfo
	r.TplName = "login/index.html"
}

func (r *IndexController) UserInfo()  {
	username := r.GetString("username")
	sql := "select game_name,role from op_giverole where user_name = ?"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql,username).ValuesList(&list)
	slice := make([]interface{},0)
	if err == nil && res > 0 {
		var gametypename string
		var gameurl string
		for i := 0;i<len(list);i++{
			roleinfo,_ := list[i][1].(string)
			arrayrole := libs.SetRole(roleinfo)
			a := fmt.Sprint(arrayrole)
			b := strings.Trim(a,"[]")
			fmt.Println(a,b)
			str := strings.Replace(strings.Trim(fmt.Sprint(arrayrole), "[]"), " ", ",", -1)
			//gameroleinfo := arrayrole.(string)
			gameAllinfo := base64.URLEncoding.EncodeToString([]byte(str))
			plattest := make(map[string]interface{})
			plattest["gameallinfo"] = gameAllinfo
			plattest["gamename"] = list[i][0]
			gamename,_ := list[i][0].(string)
			GameName,err := models.GameNameGetByGamename(gamename)
			if err != nil {
				gametypename = ""
				libs.NewLog().Error("failed",err)
			}else{
				gametypename = GameName.GamePartment
				gameurl = GameName.GameUrl
			}
			plattest["gametype"] = gametypename
			plattest["gameurl"] = gameurl
			plattest["username"] = username
			slice = append(slice,plattest)
		}
		r.Data["json"] = &slice
	}else{
		plattest := make(map[string]interface{})
		plattest["gamename"] = ""
		r.Data["json"] = &plattest
	}
	r.ServeJSON()

}

func (r *IndexController) IndexSetup(){
	name := r.GetString("name")
	User,_ := models.UsergetByusername(name)
	plattest := make(map[string]interface{})
	phonenumber := User.PhoneNumber
	email := User.EmailAdress
	plattest["phone"] = phonenumber
	plattest["email"] = email
	r.Data["json"] = &plattest
	r.ServeJSON()
}

func (r *IndexController) Prepare(){
	r.EnableXSRF = false
}


func (r *IndexController) IndexAdd(){
	username := r.GetString("username")
	User,err := models.UsergetByusername(username)
	phonenumber := r.GetString("phonenumber")
	emailadress := r.GetString("emailadress")
	_,head,err := r.GetFile("headimg")
	if phonenumber == "" || emailadress == ""{
		plattest := make(map[string]interface{})
		plattest["msg"] = "输入框不能为空"
		r.Data["json"] = plattest
		r.ServeJSON()
	}else {
		if head.Filename == ""{
			plattest := make(map[string]interface{})
			plattest["msg"] = "您还没有选择头像"
			r.Data["json"] = plattest
			r.ServeJSON()
		}else{
			if err != nil{
				libs.NewLog().Error("图片上传失败",err)
				beego.Info("文件上传失败")
			}else{
				filename := head.Filename
				//headname := strings.Split(string(filename),".")[1]
				User.PhoneNumber = phonenumber
				User.EmailAdress = emailadress
				User.HeadPortraitName = filename
				//fileExt := path.Ext(head.Filename)
				//if fileExt != ".png" && fileExt!=".jpeg" || fileExt!=".jpg"{
				//	beego.Info("文件名不正确")
				//}
				if err := User.UserUpdate();err != nil{
					libs.NewLog().Error("failed",err)
					r.ajaxMsg(err.Error(),Msg_Err)
				}else{
					err1 := r.SaveToFile("headimg","./static/img/headimg/"+filename)
					if err1 != nil{
						libs.NewLog().Error("图片保存失败",err)
						beego.Info("文件保存失败")
					}
					plattest := make(map[string]interface{})
					plattest["filename"] = filename
					plattest["msg"] = "个人设置修改成功"
					r.Data["json"] = plattest
					r.ServeJSON()
				}
			}
		}
	}
}

