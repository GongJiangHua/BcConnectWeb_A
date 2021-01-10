package controllers

import "github.com/astaxie/beego"

type BlockChain struct {
	beego.Controller
}

func (b *BlockChain)Get()  {
	b.TplName = "blockchain.html"
}