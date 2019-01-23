package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"opManage/lib"
)

type Redircontroller struct {
	beego.Controller
}

func (r *Redircontroller) Get() {
	code := r.GetString("code")
	//fmt.Println(code)
	access_token := lib.GetAccessToken()
	persistent_code,openid1 := lib.GetPerpetualCode(code,access_token)
	sns_token := lib.GetSnsToken(persistent_code,openid1,access_token)
	name,openid := lib.GetUserInfo(sns_token)
	fmt.Println(name,openid)
	if name == ""{

	}
	r.Ctx.WriteString(name)
}

