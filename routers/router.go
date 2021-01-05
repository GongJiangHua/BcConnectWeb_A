package routers

import (
	"BcConnectWeb_A/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//默认页面，默认是登录页面
    beego.Router("/", &controllers.MainController{})
    //登录接口
    beego.Router("/home",&controllers.LoginController{})
    //注册接口
    beego.Router("/register",&controllers.RegisterController{})
	//注册成功跳转到登录页面
    beego.Router("/register_sms",&controllers.RegisterSmsController{})
	//查询页面获取用户想要查询的内容，进行分析
	beego.Router("/address",&controllers.SerchResultControllers{})
	beego.Router("/blockchain",&controllers.SerchResultControllers{})
	beego.Router("/network",&controllers.SerchResultControllers{})
}
