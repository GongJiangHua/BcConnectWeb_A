package controllers

import (
	"BcConnectWeb_A/btcService"
	"fmt"
	"github.com/astaxie/beego"
)

type Two_GetWalletInfo struct {
	beego.Controller
}

func (t *Two_GetWalletInfo)Get()  {
	t.TplName = "getwalletinfo.html"
	walletinfo ,err := btcService.GetWalletInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println("最新区块的hash:", hash)
	//b.Ctx.WriteString("最新区块hash")
	t.Data["WalletInfo"]=walletinfo
}