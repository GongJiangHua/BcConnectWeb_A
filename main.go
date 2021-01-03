package main

import (
	"BtWeb/bccommand"
	"BtWeb/btcService"
	"BtWeb/db_mysql"
	_ "BtWeb/routers"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/mitchellh/mapstructure"
)

func main() {
	db_mysql.Connect()
    address:=bccommand.GetAddress()
    fmt.Println(address)
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")




	blockCount, err := btcService.GetBlockCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("区块总数：", blockCount)
	//测试：获取最新区块的hash
	bestvBlockHash, err := btcService.GetBestBlockHash()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("最新区块的hash：", bestvBlockHash)
	//测试：区块高度查询区块hash

	//测试：获取一个新地址
	newAddress, err := btcService.GetDifficulty()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("当前区块的难度：", newAddress)

	//测试：获取当前区块链信息
	blockChainInfo,err := btcService.GetBlockChainInfo()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(&blockChainInfo.Chain)


	wallinfo,err:=btcService.WallInfo()
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(wallinfo)

	mininginfo,err:=btcService.MiningInfo()
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(mininginfo)

	abortrescan,err:=btcService.Abortrescan()
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(abortrescan)

	getnettotals,err:=btcService.GetNetTotals()
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(getnettotals)



	beego.Run()
}

