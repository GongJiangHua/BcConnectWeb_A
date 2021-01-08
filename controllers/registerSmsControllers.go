package controllers

import (
	"BcConnectWeb_A/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterSmsController struct {
	beego.Controller
}

/*func (c *RegisterSmsController) Get() {
	c.Ctx.WriteString("aaaaaa")
}*/

//处理注册请求
func (c * RegisterSmsController) Post() {
	//解析
	var user models.User
	err :=c.ParseForm(&user)
	if err!= nil{
		c.TplName = "error.html"
		//c.Ctx.WriteString("抱歉，用户登录失败，请重试")
		err := "数据解析失败"
		c.Data["Error"]=err
	}
	fmt.Println("电话",user.Phone)

	//将解析的数据保存带数据库
	_,err =user.AddUser()
	if err != nil{
		fmt.Println(err.Error())
		c.TplName = "error.html"
		err := "添加用户信息失败，请重新尝试！！！"
		c.Data["Error"]=err
		//c.Ctx.WriteString("用户注册失败")
		return
	}
	//fmt.Println(res)

	c.TplName = "login.html"
}
