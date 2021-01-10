package controllers

import "github.com/astaxie/beego"

type Wallet struct {
	beego.Controller
}

func (w *Wallet) Get() {
	w.TplName = "wallet.html"
}