package controllers

import "github.com/astaxie/beego"

type LoginSmsController struct {
	beego.Controller
}

func (s *LoginSmsController) Post() {
	s.TplName = "login.html"
}