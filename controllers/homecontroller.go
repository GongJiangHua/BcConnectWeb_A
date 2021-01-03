package controllers

import "github.com/astaxie/beego"

type HoneController struct {
	beego.Controller
}

func (h *HoneController)Get()  {
	h.TplName="home.html"
}