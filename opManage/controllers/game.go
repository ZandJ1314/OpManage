package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"html/template"
	"opManage/libs"
	"opManage/models"
)

type GameController struct {
	BaseController
}

type DelGame struct {
	Name string
}

func (g *GameController) GameInfo() {
	name := g.GetString("name")
	head := g.GetString("head")
	g.Data["name"] = name
	g.Data["head"] = head
	pa,_ := g.GetInt("page")
	pre_page := 10
	totals := models.GetGameRecordNum()
	res := libs.Paginator(pa,pre_page,totals)
	res["name"] = name
	res["head"] = head
	gamelist := models.SearchGameDataList(10,pa)
	g.Data["datas"] = gamelist
	g.Data["paginator"]=res //分页的数据
	g.Data["totals"] = totals
	sql2 := "select gametypename from op_gametype;"
	o2 := orm.NewOrm()
	var list2 []orm.ParamsList
	res2,err2 := o2.Raw(sql2).ValuesList(&list2)
	if err2 == nil && res2 > 0 {
		slice2 := make([]interface{},0)
		for i := 0;i<len(list2);i++{
			plattest2 := make(map[string]interface{})
			plattest2["gametypename"] = list2[i][0]
			slice2 = append(slice2,plattest2)
		}
		g.Data["result2"] = slice2
	}else{
		libs.NewLog().Error("usertype信息提取错误",err2)
	}
	g.Data["xsrfdata"]=template.HTML(g.XSRFFormHTML())
	g.TplName = "permission/game.html"
}

func (g *GameController) Prepare(){
	g.EnableXSRF = false
}


func (g *GameController) AddGame() {
	gamename := g.GetString("gamename")
	alaisgame := g.GetString("alaisgamename")
	gameurl := g.GetString("gameurl")
	isaddtype := g.GetString("bedStatus")
	var gametype string
	if isaddtype == "allot"{
		gametype = g.GetString("addgametype")
		if gametype == ""{
			result := make(map[string]interface{})
			result["message"] = "输入框不能为空！！"
			g.Data["json"] = result
			g.ServeJSON()
		}else {
			newType := new(models.GameType)
			newType.Gametypename = gametype
			if _,err := models.GameTypeAdd(newType);err != nil{
				libs.NewLog().Error("failed",err)
			}
		}
	}else{
		gametype = g.GetString("gametype")
	}
	if gamename == "" || alaisgame == "" || gameurl == ""{
		result := make(map[string]interface{})
		result["message"] = "输入框不能为空！！"
		g.Data["json"] = result
		g.ServeJSON()
	}else{
		//newGameType,_ := models.GameTypeGetByGameTypeName(gametype)
		newGame := new(models.GameName)
		newGame.Gamename = gamename
		newGame.AlaisGamename = alaisgame
		newGame.GameUrl = gameurl
		newGame.GamePartment = gametype
		//newGame.Gametype = newGameType
		//newGame.User = User
		if _,err := models.GameNameAdd(newGame);err != nil{
			libs.NewLog().Error("failed",err)
			g.ajaxMsg(err.Error(),Msg_Err)
		}
		g.ajaxMsg("游戏增加成功",Msg_OK)
	}


}

func (g *GameController) GameDelete(){
	var delgame DelGame
	data := g.Ctx.Input.RequestBody
	err := json.Unmarshal(data,&delgame)
	if err != nil{
		libs.NewLog().Error("json.Unmarshal is err:",err.Error())
	}
	gamename := delgame.Name
	GiveRole,err := models.GiveRoleGetByGamename(gamename)
	if GiveRole == nil{
		num,err := models.GameDelete(gamename)
		if num >0 && err == nil{
			msg := gamename + "已经成功删除了！"
			g.ajaxMsg(msg,Msg_OK)
		}
	}else {
		_,err1 := models.GiveRoleDeleteByGamename(gamename)
		_,err2 := models.GameDelete(gamename)
		if err1 == nil && err2 == nil{
			msg := gamename + "已经成功删除了！"
			g.ajaxMsg(msg,Msg_OK)
		}
	}
}

func (g *GameController) GameDetail() {
	var detgame DelGame
	data := g.Ctx.Input.RequestBody
	err := json.Unmarshal(data,&detgame)
	if err != nil{
		libs.NewLog().Error("json.Unmarshal is err:",err.Error())
	}
	gamename := detgame.Name
	sql := "select user_name from op_giverole where game_name = ?;"
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql,gamename).ValuesList(&list)
	slice := make([]interface{},0)
	if err == nil && res > 0{
		fmt.Println("hello")
		for i := 0;i<len(list);i++{
			plattest := make(map[string]interface{})
			plattest["username"] = list[i][0]
			slice = append(slice,plattest)
		}
		g.Data["json"] = slice
	}else{
		plattest := make(map[string]interface{})
		plattest["username"] = "该游戏暂无管理员管理！"
		g.Data["json"] = plattest
		libs.NewLog().Error("gamename信息提取错误",err)
	}
	g.ServeJSON()
}

func (g *GameController) JudgeGame() {
	name := g.GetString("name")
	_,err := models.GameNameGetByGamename(name)
	plattest := make(map[string]interface{})
	if err == nil{
		plattest["name"] = "false"
	}else{
		plattest["name"] = "true"
	}
	g.Data["json"] = &plattest
	g.ServeJSON()
}

func (g *GameController) JudgeType() {
	name := g.GetString("name")
	_,err := models.GameTypeGetByGameTypeName(name)
	plattest := make(map[string]interface{})
	if err == nil{
		plattest["name"] = "false"
	}else{
		plattest["name"] = "true"
	}
	g.Data["json"] = &plattest
	g.ServeJSON()
}

func (g *GameController) GameSetupInfo(){
	name := g.GetString("name")
	GameName,_ := models.GameNameGetByGamename(name)
	plattest := make(map[string]interface{})
	alaisgamename := GameName.AlaisGamename
	gameurl := GameName.GameUrl
	gametype := GameName.GamePartment
	plattest["alaisgamename"] = alaisgamename
	plattest["gameurl"] = gameurl
	plattest["gametype"] = gametype
	g.Data["json"] = &plattest
	g.ServeJSON()

}

func (g *GameController) GameUpdate() {
	gamename := g.GetString("gamenamechange")
	GameName,_ := models.GameNameGetByGamename(gamename)
	alaisgame := g.GetString("alaisgamenamechange")
	gameurl := g.GetString("gameurlchange")
	gametype := g.GetString("changegametype")
	_,err := models.GameTypeGetByGameTypeName(gametype)
	if err != nil{
		newType := new(models.GameType)
		newType.Gametypename = gametype
		if _,err := models.GameTypeAdd(newType);err != nil{
			libs.NewLog().Error("failed",err)
		}
	}
	if alaisgame == "" || gameurl == "" || gametype == ""{
		result := make(map[string]interface{})
		result["message"] = "输入框不能为空！！"
		g.Data["json"] = result
		g.ServeJSON()
	}else{
		GameName.AlaisGamename = alaisgame
		GameName.GameUrl = gameurl
		GameName.GamePartment = gametype
		if err1 := GameName.GamenameUpdate();err1 != nil{
			libs.NewLog().Error("failed",err)
			g.ajaxMsg(err.Error(),Msg_Err)
		}else{
			g.ajaxMsg("游戏修改成功",Msg_OK)
		}
	}
}

















