package controllers

import (
	"BcConnectWeb_A/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginSmsController struct {
	beego.Controller
}

func (l *LoginSmsController) Get() {
	l.TplName = "login_sms.html"
}

/**
短信验证码登录功能
 */
func (l *LoginSmsController) Post() {
	var smsLogin models.SmsLogin

	err := l.ParseForm(&smsLogin)
	if err != nil {
		//l.Ctx.WriteString("抱歉，验证码登陆失败，请重试")
		l.Data["Error"] = "抱歉，验证码登陆失败，请重试"
		l.TplName = "err.html"
		return
	}
	//先验证手机号是否注册
	user := models.User{
		Phone:    smsLogin.Phone,
	}
	u,err := user.QuerUser()
	if err != nil {
		//l.Ctx.WriteString("对不起，登陆失败，请重试")
		l.Data["Error"] = "对不起，登陆失败，请重试"
		l.TplName = "err.html"
		return
	}
	if u.Phone == "" {
		l.Data["Error"] = "对不起，您还为注册，请先注册"
		l.TplName = "err.html"
		return
	}

	//查询用户提交的登录信息是否正确
	sms,err := models.QuerySmsRecord(smsLogin.Phone,smsLogin.Code)

	if err != nil {
		//l.Ctx.WriteString("抱歉")
		l.Data["Error"] = err.Error()
		l.TplName = "err.html"
		fmt.Println("错了：",err.Error())
	}

	if sms.BizId == "" {//验证码错误，手机号错误
		//l.Ctx.WriteString("手机号或验证码错误，请重新输入")
		l.Data["Error"] = "手机号或验证码错误，请重新输入"
		l.TplName = "err.html"
		return
	}
	//now := time.Now().Unix()
	//if (now - sms.TimeStamp) > 30000 {
	//	l.Ctx.WriteString("验证码已失效，请重新获取")
	//	return
	//}

	l.TplName = "zhihu_login_sms.html"
}
