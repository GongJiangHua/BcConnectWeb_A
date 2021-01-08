package controllers

import (
	"BcConnectWeb_A/btcService"
	"fmt"
	"github.com/astaxie/beego"
)

type Three_GetDifficulty struct {
	beego.Controller
}

func (s *Three_GetDifficulty) Get()  {
	s.TplName = "getdifficulty.html"
	difficulty ,err := btcService.GetDifficulty()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println("最新区块的hash:", hash)
	//b.Ctx.WriteString("最新区块hash")
	s.Data["Difficulty"]=difficulty
}
