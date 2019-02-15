package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"opManage/lib"
	"opManage/models"
)

type IndexController struct {
	BaseController
}


func (r *IndexController) Get() {
	code := r.GetString("code")
	//fmt.Println(code)
	access_token := lib.GetAccessToken()
	persistent_code,openid1 := lib.GetPerpetualCode(code,access_token)
	sns_token := lib.GetSnsToken(persistent_code,openid1,access_token)
	if sns_token == "errmsg"{
		//fmt.Println(sns_token,"hehehehehehehheheheh")
		r.TplName = "login/index.html"
	}
	name,openid := lib.GetUserInfo(sns_token)
	if name == "error"{
		r.TplName = "login/index.html"
	}
	fmt.Println(name,openid)
	dataurl := r.Ctx.Request.RequestURI
	lens := len(dataurl)
	//fmt.Println(dataurl[lens-5:])
	//a := reflect.TypeOf(dataurl).String()//判断数据类型
	//fmt.Println(a)
	userinfo := make(map[string]interface{})
	if dataurl[lens-5:] == "LOGIN"{
		User,err := models.UsergetByOpenid(openid)
		fmt.Println(err,User)
		if err != nil{
			//r.ajaxMsg(err.Error(),Msg_Err)
			//r.Data["result"] = "您还未进行注册，请点击下面链接进行注册"
			r.redirect("/login")
		}else{
			id := User.Id
			userinfo["id"] = id
			name := User.UserName
			userinfo["name"] = name
			openid := User.Openid
			userinfo["openid"] = openid
			headPortraitName := User.HeadPortraitName
			userinfo["headPortraitName"] =headPortraitName
			manageName := User.ManageName
			userinfo["manageName"] = manageName
			issuperadministrator := User.Issuperadministrator
			userinfo["issuperadministrator"] = issuperadministrator
			r.Data["result"] = userinfo
			r.TplName = "login/index.html"
		}
	}else{
		_,err := models.UsergetByOpenid(openid)
		if err != nil{
			if openid != "error"{
				newUser := new(models.User)
				newUser.UserName = name
				newUser.Openid = openid
				if _,err := models.UserAdd(newUser);err != nil {
					r.ajaxMsg(err.Error(),Msg_Err)
				}
				User,err:= models.UsergetByOpenid(openid)
				if err != nil{
					lib.NewLog().Error("failed",err)
				}
				//fmt.Println(User,err,"hello world")
				id := User.Id
				userinfo["id"] = id
				name := User.UserName
				userinfo["name"] = name
				openid := User.Openid
				userinfo["openid"] = openid
				headPortraitName := User.HeadPortraitName
				userinfo["headPortraitName"] =headPortraitName
				manageName := User.ManageName
				userinfo["manageName"] = manageName
				issuperadministrator := User.Issuperadministrator
				userinfo["issuperadministrator"] = issuperadministrator
				r.Data["result"] = userinfo
				r.TplName = "login/index.html"
			}else{
				lib.NewLog().Error("failed,openid传值进来为%s",openid)
			}
		}else{
			//r.Data["result"] = "用户已经存在,请点击下面链接进行登录"
			r.redirect("/register")
		}

	}
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
			fmt.Println(head.Filename,err)
			fmt.Println(phonenumber,emailadress)
			r.TplName = "login/index.html"
		}
	}

}
