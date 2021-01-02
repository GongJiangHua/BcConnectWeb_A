package controllers

import (
	"BcConnectWeb_A/models"
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
	fmt.Println("order：",order)
	fmt.Println(order.Order)


	q.Ctx.WriteString("查到了")
}
