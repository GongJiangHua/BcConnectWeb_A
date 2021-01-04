package btcService

import (
	models2 "BcConnectWeb_A/models"
	"BcRPCCode/entity"
	"BtWeb/models"
	"BtWeb/utils"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	_ "github.com/mitchellh/mapstructure"
)

//该包用来封装bitcoin命令

//用于查询区块链中的区块数量
func GetBlockCount() (interface{}, error) {
	jsonMes, err := utils.PrepareJsonStr("getblockcount", nil)
	if err != nil {
		fmt.Println("准备json传输数据失败，错误原因：", err.Error())
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		fmt.Println("请求RPC服务失败，错误原因：", err.Error())
		return nil, err
	}
	Rpcresult :=models.RPCResult{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil
}

//用于查询区块链中最新区块的hash

func GetBestBlockHash() (interface{}, error) {
	jsonMes, err := utils.PrepareJsonStr("getbestblockhash", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := models.RPCResult{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil
}


func WallInfo() (*models.WallInfo, error) {
	jsonMes, err := utils.PrepareJsonStr("getwalletinfo", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := models.WallInfo{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	var walletinfo models.WallInfo
	err = mapstructure.Decode(walletinfo,&Rpcresult)
	if err!=nil {
		return nil,err
	}
	return &walletinfo, nil
}


func WallInfo() (*models.WallInfo, error) {
	jsonMes, err := utils.PrepareJsonStr("getwalletinfo", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := models.WallInfo{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	var walletinfo models.WallInfo
	err = mapstructure.Decode(walletinfo,&Rpcresult)
	if err!=nil {
		return nil,err
	}
	return &walletinfo, nil
}




func RpcInfo() (*models2.RpcInfo, error) {
	jsonMes, err := utils.PrepareJsonStr("getmininginfo", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := models.MiningInfo{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	var rpcInfo models2.RpcInfo
	err = mapstructure.Decode(rpcInfo,&Rpcresult)
	if err!=nil {
		return nil,err
	}
	return &rpcInfo, nil
}

func Abortrescan() (*models.Abortrescan, error) {
	jsonMes, err := utils.PrepareJsonStr("abortrescan", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := models.Abortrescan{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	var abortrescans models.Abortrescan
	err = mapstructure.Decode(abortrescans,&Rpcresult)
	if err!=nil {
		return nil,err
	}
	return &abortrescans, nil
}

func GetNetTotals() (*models.GetNetToTals, error) {
	jsonMes, err := utils.PrepareJsonStr("getnettotals", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := models.GetNetToTals{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	var getnettotals models.GetNetToTals
	err = mapstructure.Decode(getnettotals,&Rpcresult.Uploadtarget)
	if err!=nil {
		return nil,err
	}
	return &getnettotals, nil
}

//根据区块高度查询区块hash
//func GetBlockHashByHeight() (interface{}, error) {
//	jsonMes, err := utils.PrepareJsonStr("getblockhash",nil)
//	if err != nil {
//		return nil, err
//	}
//	resBytes, err := utils.SendRPCPost(jsonMes)
//	if err != nil {
//		return nil, err
//	}
//	Rpcresult := entity.RPCResult{}
//	err = json.Unmarshal(resBytes,&Rpcresult)
//	if err != nil {
//		return nil,err
//	}
//	return Rpcresult.Result, nil
//}

//用于获取当前区块的难度
func GetDifficulty() (interface{}, error) {
	jsonMes, err := utils.PrepareJsonStr("getdifficulty", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := entity.RPCResult{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil
}

//用于获取当前区块链信息
func GetBlockChainInfo() (*models.BlockChainInfo,error) {
	jsonMes, err := utils.PrepareJsonStr("getblockchaininfo",nil)
	if err != nil {
		return nil,err
	}
	resBytes,err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil,err
	}
	Rpcresult := models.BlockChainInfo{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return &Rpcresult,nil
}
