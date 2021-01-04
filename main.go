package BcConnectWeb_A

import (
	"BcConnectWeb_A/btcService"
	"fmt"
	"github.com/astaxie/beego"
)

func main()  {


	result,err:=btcService.RpcInfo()
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result.Active_commands)


	getnet,err:=btcService.GetNetTotals()
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(getnet)

	beego.Run()
}
