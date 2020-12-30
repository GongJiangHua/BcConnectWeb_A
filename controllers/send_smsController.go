package controllers

import (
	"BcConnectWeb_A/models"
	"BcConnectWeb_A/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type SendSmsController struct {
	beego.Controller
}

func (s *SendSmsController) Get() {
	s.TplName = "send_sms.html"
}

/**
发送验证码功能
 */
func (s *SendSmsController) Post() {
	var smslogin models.SmsLogin
	err := s.ParseForm(&smslogin)
	if err != nil {
		s.Ctx.WriteString("抱歉，验证码获取失败，请重试")
		return
	}

	phone := smslogin.Phone
	code := utils.GenRandCode(6)//返回一个六位数的随机数
	result,err := utils.SemdSms(phone,code,utils.SMS_TLP_REGISTER)
	if err != nil {
		s.Ctx.WriteString("验证码发送失败，请重试")
	}
	if len(result.BizId) == 0 {
		s.Ctx.WriteString("错误："+result.Message)
		return
	}
	//到此，验证码发送成功
	smsRecord := models.SmsRecord{
		BizId:     result.BizId,
		Phone:     phone,
		Code:      code,
		Status:    result.Code,
		Message:   result.Message,
		TimeStamp: time.Now().Unix(),
	}
	_,err = smsRecord.SaveSmsRecord()
	if err != nil {
		s.Ctx.WriteString("抱歉，验证码说去失败，请重试")
		fmt.Println("有错误：",err.Error())
		return
	}

	//验证记录保存成功


	s.TplName = "login_sms.html"
}