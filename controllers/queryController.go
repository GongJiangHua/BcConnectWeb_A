package controllers

import (
	"BcConnectWeb_A/models"
	"BcConnectWeb_A/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type QueryController struct {
	beego.Controller
}

func (q QueryController) Post() {
	var order models.QueryOrder
	err := q.ParseForm(&order)
	if err != nil {
		q.Ctx.WriteString("抱歉，数据解析失败")
		return
	}

	Btc := utils.Conf()
	fmt.Println("len(btc)",len(Btc))
	for key,value :=range Btc{
		fmt.Println("key,value:",key,value)
	}

	q.Ctx.WriteString("查到了")
}
