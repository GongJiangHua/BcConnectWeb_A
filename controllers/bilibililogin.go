package controllers

import "github.com/astaxie/beego"

type BilibiliLogin struct {
	beego.Controller

}

func (b *BilibiliLogin)Get()  {
	b.TplName="bilibiliregister.html"
}