package controllers


type GameController struct {
	BaseController
}

func (g *GameController) Get() {
	g.TplName = "permission/game.html"
}
