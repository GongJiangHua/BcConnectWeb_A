package controllers

import "github.com/astaxie/beego"

type NetWork struct {
	beego.Controller
}

func (n *NetWork)Get()  {
	n.TplName = "network.html"
}