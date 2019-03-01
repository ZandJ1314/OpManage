package controllers


type SystemController struct {
	BaseController
}

func (s *SystemController) Get(){
	name := s.GetString("name")
	head := s.GetString("head")
	s.Data["name"] = name
	s.Data["head"] = head
	s.TplName = "permission/system.html"
}
