package routers

import (
	"BcConnectWeb_A/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //登录接口
    beego.Router("/home",&controllers.LoginController{})

    //注册接口
    beego.Router("/register",&controllers.RegisterController{})


   beego.Router("/register_sms",&controllers.RegisterSmsController{})

    //验证码登录接口
    beego.Router("/login_sms",&controllers.LoginSmsController{})
    //发送验证码接口
    beego.Router("/send_sms",&controllers.SendSmsController{})
    //查询btc信息接口
    beego.Router("/query",&controllers.QueryController{})


}
