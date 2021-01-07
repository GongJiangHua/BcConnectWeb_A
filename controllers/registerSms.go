package controllers

import "github.com/astaxie/beego"

type RegisterSms struct {
	beego.Controller
}

func (r *RegisterSms) Post() {
	r.TplName = "login.html"
}