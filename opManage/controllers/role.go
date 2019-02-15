package controllers

import "fmt"

type RoleController struct {
	BaseController
}

func (r *RoleController) Get(){
	sql := "select user_name from op_user;"
	fmt.Println(sql)
}
