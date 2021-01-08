package controllers

import (
	"BcConnectWeb_A/btcService"
	"fmt"
	"github.com/astaxie/beego"
)

type One_GetNewAddress struct {
	beego.Controller
}

func (o *One_GetNewAddress)Get()  {
	o.TplName = "getnewaddress.html"
	newaddress ,err := btcService.GetNewAddress()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println("最新区块的hash:", hash)
	//b.Ctx.WriteString("最新区块hash")
	o.Data["NewAddress"]=newaddress
}
