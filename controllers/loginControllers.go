package controllers

import (
	"BcConnectWeb_A/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "home.html"
}
func (c *LoginController) Post() {
	//1、解析客户端用户提交的登录数据
	var user models.User
	err := c.ParseForm(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.TplName = "error.html"
		err := "抱歉，用户登录信息解析失败，请重试!!!"
		c.Data["Error"]=err
		return
	}
	fmt.Println("user、password:",user.Phone,user.Password)
	//2、根据解析到的数据,执行数据库查询操作
	_, err = user.QuerUser()
	//3、判断数据库查询结果
	if err != nil {
		//fmt.Println(err.Error())
		//fmt.Println("ffff为:",u)
		//fmt.Println("2222",err.Error())
		c.TplName = "error.html"
		//c.Ctx.WriteString("抱歉，用户登录失败，请重试")
		err := "抱歉，用户登录失败，请检查你的用户名和密码，重新尝试!!!"
		c.Data["Error"]=err
		return
	}
	fmt.Println("user、password:",user.Phone,user.Password)
	c.TplName = "home.html"
}

