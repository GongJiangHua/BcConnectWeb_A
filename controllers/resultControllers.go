package controllers

import (
	"BcConnectWeb_A/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type SerchResultControllers struct {
	beego.Controller
}

func (s *SerchResultControllers) Post() {
	fmt.Println("正在运行Search查询功能")
	var serchText models.SearchMes
	err := s.ParseForm(&serchText)
	if err != nil {
		s.Ctx.WriteString("获取用户查询信息失败，请重试")
		return
	}
	fmt.Println("用户想要查询的内容",serchText)
	fmt.Println("用户想要查询的内容",serchText.SearchText)
	if strings.ToUpper(serchText.SearchText)=="BLOCKCHAIN" {
		s.TplName = "blockchain.html"
	}else if strings.ToUpper(serchText.SearchText)=="WALLET" {
		s.TplName = "wallet.html"
	}else if strings.ToUpper(serchText.SearchText)=="NETWORK" {
		s.TplName = "network.html"
	}else {
		s.TplName = "error.html"
		err := "请输入正确的查询范围！！！"
		s.Data["Error"]=err
		//s.Ctx.WriteString("请输入正确的查询范围！！！")
	}
	//s.Data["SearchMes"]=serchText
}
