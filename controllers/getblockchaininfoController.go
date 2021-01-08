package controllers

import (
	"BcConnectWeb_A/btcService"
	"fmt"
	"github.com/astaxie/beego"
)

type Two_BlockChainInfoControllers struct {
	beego.Controller
}

func (t *Two_BlockChainInfoControllers) Get()  {
	t.TplName = "getblockchaininfo.html"
	BlockChainInfo ,err := btcService.GetBlockChainInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("BlockChainInfo:",BlockChainInfo.Blocks)
	//fmt.Println("最新区块的hash:", hash)
	//b.Ctx.WriteString("最新区块hash")
	t.Data["BlockChainInfo"]=BlockChainInfo
}