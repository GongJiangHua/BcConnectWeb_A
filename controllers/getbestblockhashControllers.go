package controllers

import (
	"BcConnectWeb_A/btcService"
	"github.com/astaxie/beego"
)

type One_bestblockhashControllers struct {
	beego.Controller
}

func (o *One_bestblockhashControllers) Get()  {
	o.TplName = "getbestblockhash.html"
	hash, _:= btcService.GetBestBlockHash()
	//fmt.Println("最新区块的hash:", hash)
	//b.Ctx.WriteString("最新区块hash")
	text := "BestBlockHash:"
	o.Data["Text"] = text
	o.Data["BestBlockHash"]=hash
}