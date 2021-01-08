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
    //beego.Router("/homeone.html",&controllers.LoginController{})
    //注册接口
    beego.Router("/register",&controllers.RegisterController{})
	//注册成功跳转到登录页面
    beego.Router("/register_sms",&controllers.RegisterSmsController{})
	//查询页面获取用户想要查询的内容，进行分析
	beego.Router("/index",&controllers.SerchResultControllers{})
    //SIGN IN实现用户直接登录
    beego.Router("/login.html",&controllers.MainController{})
	//有关比特币页面显示的路由
	beego.Router("/getbestblockhash.html",&controllers.One_bestblockhashControllers{})
	beego.Router("/getblockchaininfo.html",&controllers.Two_BlockChainInfoControllers{})
	beego.Router("/getdifficulty.html",&controllers.Three_GetDifficulty{})
	beego.Router("/getblockcount.html",&controllers.Four_GetBlockCount{})
	beego.Router("/getnewaddress.html",&controllers.One_GetNewAddress{})
	beego.Router("/getwalletinfo.html",&controllers.Two_GetWalletInfo{})
    //验证码登录接口
	beego.Router("/login_sms",&controllers.LoginSmsController{})
	//发送验证码接口
	beego.Router("/send_sms",&controllers.SendSmsController{})
	//beego.Router("/blockchain.html",&controllers.BlockChainControllers{})
	//beego.Router("/blockchain.html",&controllers.BlockChainControllers{})
	//beego.Router("/blockchain.html",&controllers.BlockChainControllers{})
	//beego.Router("/blockchain.html",&controllers.BlockChainControllers{})
}
