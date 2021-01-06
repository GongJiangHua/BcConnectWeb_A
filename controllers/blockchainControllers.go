package controllers

import (
	"BcConnectWeb_A/btcService"
	"BcConnectWeb_A/entity"
	"github.com/astaxie/beego"
)

type BlockChainControllers struct {
	beego.Controller
}

func (b *BlockChainControllers) Get()  {
	
	var bestblockhash entity.BestBlockHash
	hash, _:= btcService.GetBestBlockHash()
	bestblockhash.Bestblockhash=hash
	//fmt.Println("最新区块的hash:", hash)
	b.Data["BestBlockHash"]=bestblockhash
	b.TplName = "blockchain.html"
}