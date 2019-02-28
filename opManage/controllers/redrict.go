package controllers

import (
	"fmt"
	"opManage/lib"
	"opManage/models"
)

type Redircontroller struct {
	BaseController
}


func (r *Redircontroller) Get() {
	code := r.GetString("code")
	//fmt.Println(code)
	access_token := lib.GetAccessToken()
	persistent_code,openid1 := lib.GetPerpetualCode(code,access_token)
	sns_token := lib.GetSnsToken(persistent_code,openid1,access_token)
	name,openid := lib.GetUserInfo(sns_token)
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
			r.Data["result"] = "您还未进行注册"
			r.TplName = "login/login.html"
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
			UserinfoControll(userinfo)
		}
	}else{
		_,err := models.UsergetByOpenid(openid)
		if err != nil{
			newUser := new(models.User)
			newUser.UserName = name
			newUser.Openid = openid
			newUser.HeadPortraitName = name
			if _,err := models.UserAdd(newUser);err != nil {
				r.ajaxMsg(err.Error(),Msg_Err)
			}
			User,_ := models.UsergetByOpenid(openid)
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
			UserinfoControll(userinfo)
		}
		r.Data["result"] = "用户已经存在,请直接登录"
		r.TplName = "register/register.html"

	}
}

