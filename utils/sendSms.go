package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"time"
)

type SmsCode struct {
	Code string`json:"code"`
}

type SmsResult struct {
	BizId string
	Code string
	Message string
	RequestId string
}

const SMS_TLP_REGISTER = "SMS_176525619"//注册业务的短信模板
const SMS_TLP_LOGIN  = "SMS_176525619"//用户登陆的短信模板
const SMS_TLP_KYC  = "SMS_176525619"//实名认证的短信模板

/**
用于发送一条短信息
 */
func SemdSms(phone string,code string,template string) (*SmsResult,error) {
	config := beego.AppConfig
	//获取配置文件中的数据
	smsAccesKey := config.String("sms_acces_key")
	smsAccesSecret := config.String("sms_acces_secret")

	client,err := dysmsapi.NewClientWithAccessKey("cn-hangzhou",smsAccesKey,smsAccesSecret)
	if err != nil {
		return nil,err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = phone//要发送的手机号
	request.SignName = "线上餐厅"
	request.TemplateCode = template//指定短信模板
	smsCode := SmsCode{
		Code:code,
	}

	smsbytes,_ := json.Marshal(smsCode)
	request.TemplateParam = string(smsbytes)

	response,err := client.SendSms(request)
	if err != nil {
		fmt.Println("这：",err.Error())
	}
	smsResule := &SmsResult{
		BizId:     response.BizId,
		Code:      response.Code,
		Message:   response.Message,
		RequestId: response.RequestId,
	}
	return smsResule,nil

}

func GenRandCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i:=0;i<width ; i++ {
		fmt.Fprintf(&sb,"%d",numeric[rand.Intn(r)])
	}
	return sb.String()
}
