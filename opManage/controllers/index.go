package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/gomodule/redigo/redis"
	"html/template"
	"opManage/lib"
	"opManage/models"
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
	GameInfo interface{}
}


func (r *IndexController) Get() {
	name := r.Ctx.GetCookie("username")
	sessionid := r.Ctx.GetCookie("sessionid")
	redis,err := cache.NewCache("redis",`{"conn":"127.0.0.1:6379", "key":"cacheuserinfo"}`)
	if name == "" && sessionid == ""{
		fmt.Println("first")
		code := r.GetString("code")
		access_token := lib.GetAccessToken()
		persistent_code,openid1 := lib.GetPerpetualCode(code,access_token)
		sns_token := lib.GetSnsToken(persistent_code,openid1,access_token)
		if sns_token == "errmsg"{
			r.TplName = "login/login.html"
		}
		name,openid := lib.GetUserInfo(sns_token)
		if name == "error"{
			r.TplName = "login/index.html"
		}
		r.Ctx.SetCookie("username",name,100)
		r.Ctx.SetCookie("sessionid",openid,100)
		userinfo := make(map[string]interface{})
		//if dataurl[lens-5:] == "LOGIN"{
		User,err := models.UsergetByusername(name)
		if err != nil{
			lib.NewLog().Error("failed",err)
			r.Data["result"] = "您属于非法用户，请与管理员确认您的身份信息"
			r.TplName = "login/login.html"
		}else{
			User.Openid = openid
			if err := User.UserUpdate();err != nil{
				lib.NewLog().Error("failed",err)
			}
			name := User.UserName
			sql := "select game_name from op_giverole where user_name = ?"
			o := orm.NewOrm()
			var list []orm.ParamsList
			res,err := o.Raw(sql,name).ValuesList(&list)
			slice := make([]interface{},0)
			if err == nil && res > 0 {
				var gametypename string
				for i := 0;i<len(list);i++{
					plattest := make(map[string]interface{})
					plattest["gamename"] = list[i][0]
					gamename,_ := list[i][0].(string)
					GameName,err := models.GameNameGetByGamename(gamename)
					if err != nil {
						gametypename = ""
						lib.NewLog().Error("failed",err)
					}else{
						gametypename = GameName.GamePartment
					}
					plattest["gametype"] = gametypename
					slice = append(slice,plattest)
					fmt.Println(slice)
				}
			}
			userinfo["gameinfo"] = slice
			userinfo["name"] = name
			headPortraitName := User.HeadPortraitName
			userinfo["headPortraitName"] =headPortraitName
			manageName := User.ManageName
			userinfo["managename"] = manageName
			issuperadministrator := User.Issuperadministrator
			userinfo["issuperadministrator"] = issuperadministrator
			if err != nil{
				lib.NewLog().Error("failed","缓存数据添加错误")
			}
			//将map类型转换为json字符串，然后存入在redis中，后面直接转换就行
			value,_ := json.Marshal(userinfo)
			redis.Put("sessionid",openid,time.Second*100)
			redis.Put("userinfo",value,time.Second*100)
			r.Data["xsrfdata"]=template.HTML(r.XSRFFormHTML())
			r.Data["result"] = userinfo
			r.TplName = "login/index.html"
		}
	}else{
		//redis,err := cache.NewCache("redis",`{"conn":"127.0.0.1:6379", "key":"cacheuserinfo"}`)
		if err != nil{
			lib.NewLog().Error("failed","缓存数据存取错误")
		}
		sessionid := redis.Get("sessionid")
		userinfo := redis.Get("userinfo")
		sessionid = string(sessionid.([]byte))
		//a := reflect.TypeOf(userinfo).String()//判断数据类型
		var users Users
		err := json.Unmarshal(userinfo.([]byte),&users)
		if err != nil{
			lib.NewLog().Error("redis.failed",err)
		}
		newuserinfo := make(map[string]interface{})
		newuserinfo["name"] = users.Name
		newuserinfo["headPortraitName"] = users.HeadPortraitName
		newuserinfo["managename"] = users.ManageName
		newuserinfo["issuperadministrator"] = users.Issuperadministrator
		newuserinfo["gameinfo"] = users.GameInfo
		r.Data["xsrfdata"]=template.HTML(r.XSRFFormHTML())
		r.Data["result"] = newuserinfo
		r.TplName = "login/index.html"
	}

	//fmt.Println(openid)
	//fmt.Println(name,openid)
	//dataurl := r.Ctx.Request.RequestURI
	//lens := len(dataurl)
	//fmt.Println(dataurl[lens-5:])
	//a := reflect.TypeOf(dataurl).String()//判断数据类型
	//fmt.Println(a)

	//}else{
	//	_,err := models.UsergetByOpenid(openid)
	//	if err != nil{
	//		if openid != "error"{
	//			newUser := new(models.User)
	//			newUser.UserName = name
	//			newUser.Openid = openid
	//			if _,err := models.UserAdd(newUser);err != nil {
	//				r.ajaxMsg(err.Error(),Msg_Err)
	//			}
	//			User,err:= models.UsergetByOpenid(openid)
	//			if err != nil{
	//				lib.NewLog().Error("failed",err)
	//			}
	//			//fmt.Println(User,err,"hello world")
	//			id := User.Id
	//			userinfo["id"] = id
	//			name := User.UserName
	//			userinfo["name"] = name
	//			openid := User.Openid
	//			userinfo["openid"] = openid
	//			headPortraitName := User.HeadPortraitName
	//			userinfo["headPortraitName"] =headPortraitName
	//			manageName := User.ManageName
	//			userinfo["manageName"] = manageName
	//			issuperadministrator := User.Issuperadministrator
	//			userinfo["issuperadministrator"] = issuperadministrator
	//			r.Data["result"] = userinfo
	//			r.TplName = "login/index.html"
	//		}else{
	//			lib.NewLog().Error("failed,openid传值进来为%s",openid)
	//		}
	//	}else{
	//		//r.Data["result"] = "用户已经存在,请点击下面链接进行登录"
	//		r.redirect("/register")
	//	}
	//
	//}
}


func (r *IndexController) Post(){
	username := r.GetString("username")
	User,err := models.UsergetByusername(username)
	phonenumber := r.GetString("phonenumber")
	emailadress := r.GetString("emailadress")
	_,head,err := r.GetFile("headimg")
	fmt.Println(username,phonenumber,emailadress)
	fmt.Println(err)
	if err != nil{
		lib.NewLog().Error("图片上传失败",err)
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
			lib.NewLog().Error("failed",err)
			r.ajaxMsg(err.Error(),Msg_Err)
		}else{
			err1 := r.SaveToFile("headimg","./static/img/headimg/"+filename)
			if err1 != nil{
				lib.NewLog().Error("图片保存失败",err)
				beego.Info("文件保存失败")
			}
			//fmt.Println(head.Filename,err)
			//fmt.Println(phonenumber,emailadress)
			r.redirect("/")
		}
	}

}
