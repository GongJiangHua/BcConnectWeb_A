package controllers

import (
	"BcConnectWeb_A/btcService"
	"fmt"
	"github.com/astaxie/beego"
)

type Four_GetBlockCount struct {
	beego.Controller
}

func (t *Four_GetBlockCount) Get (){
	t.TplName = "getdifficulty.html"
	blockcount ,err := btcService.GetBlockCount()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println("最新区块的hash:", hash)
	//b.Ctx.WriteString("最新区块hash")
	t.Data["BlockCount"]=blockcount
}

