package controllers

import (
	"opManage/libs"
	"opManage/models"
)

type LogsController struct {
	BaseController
}

func (self *LogsController) LogsInfo()  {
	name := self.GetString("name")
	head := self.GetString("head")
	self.Data["name"] = name
	self.Data["head"] = head
	pa,_ := self.GetInt("page")
	pre_page := 10
	totals := models.GetLogRecordNum()
	res := libs.Paginator(pa,pre_page,totals)
	res["name"] = name
	res["head"] = head
	logslist := models.SearchLogDataList(10,pa)
	self.Data["datas"] = logslist
	self.Data["paginator"] = res
	self.Data["totals"] = totals
	self.TplName = "permission/logs.html"
}
