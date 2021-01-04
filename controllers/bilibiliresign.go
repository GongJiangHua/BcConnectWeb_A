package controllers

import "github.com/astaxie/beego"

type BilibiliResign struct {
	beego.Controller

}

func (b *BilibiliResign)Get()  {
	b.TplName="bilibiliregister.html"
}