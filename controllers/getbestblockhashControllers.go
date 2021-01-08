package controllers

import (
	"BcConnectWeb_A/btcService"
	"github.com/astaxie/beego"
)

type BlockChainControllers struct {
	beego.Controller
}

func (b *BlockChainControllers) Get()  {
	b.TplName = "getbestblockhash.html"
	hash, _:= btcService.GetBestBlockHash()
	//fmt.Println("最新区块的hash:", hash)
	//b.Ctx.WriteString("最新区块hash")
	b.Data["BestBlockHash"]=hash
}