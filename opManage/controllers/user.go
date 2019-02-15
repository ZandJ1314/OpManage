package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"opManage/lib"
)

type UserController struct {
	BaseController
}

func (u *UserController) Get(){
	sql := "select openid,user_name,phone_number,email_adress,issuperadministrator,create_time from op_user;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql).ValuesList(&list)
	if err == nil && res > 0{
		slice := make([]interface{},0)
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["openid"] = list[i][0]
			plattest["username"] = list[i][1]
			plattest["phonenumber"] = list[i][2]
			plattest["emailadress"] = list[i][3]
			plattest["issuperadministrator"] = list[i][4]
			plattest["createtime"] = list[i][5]
			fmt.Println(plattest)
			slice = append(slice,plattest)
			fmt.Println(slice)
		}
		u.Data["result"] = slice
		u.TplName = "permission/user.html"

	}else{
		lib.NewLog().Error("user信息提取错误",err)
	}
}
