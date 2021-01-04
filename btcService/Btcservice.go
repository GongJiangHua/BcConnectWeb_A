package btcService

import (
	"BcConnectWeb_A/utils"
	"fmt"
)

//该包用来封装bitcoin命令

//用于查询区块链中的区块数量
func GetBlockCount() (interface{}, error) {
	Rpcresult,err := utils.GetMsgByCommand("getblockcount",nil)
	if err != nil {
		fmt.Println("请求RPC服务失败，错误原因：", err.Error())
		return nil, err
	}
	return Rpcresult.Result, nil
}

//用于查询区块链中最新区块的hash

//func GetBestBlockHash() (interface{}, error) {
//	Rpcresult, err := utils.GetMsgByCommand("getbestblockhash",nil)
//	if err != nil {
//		return nil, err
//	}
//	return Rpcresult.Result, nil
//}

//根据区块高度查询区块hash
//func GetBlockHashByHeight(height int) (interface{}, error) {
//	Rpcresult, err := utils.GetMsgByCommand("getblockhash",height)
//	if err != nil {
//		return nil, err
//	}
//	return Rpcresult.Result, nil
//}

//用于获取当前区块的难度
//func GetDifficulty() (interface{}, error) {
//	Rpcresult, err := utils.GetMsgByCommand("getdifficulty",nil)
//	if err != nil {
//		return nil, err
//	}
//	return Rpcresult.Result, nil
//}

//用于获取当前区块链信息
//func GetBlockChainInfo() (*entity.BlockChainInfo,error) {
//	Rpcresult,err := utils.GetMsgByCommand("getblockchaininfo",nil)
//	if err != nil {
//		return nil,err
//	}
//	var blockchaininfo entity.BlockChainInfo
//	err = mapstructure.Decode(Rpcresult,&blockchaininfo)
//	return &blockchaininfo,nil
//}

////生成一个新地址
//func GetNewAddress()(interface{},error)  {
//	Rpcresult, err := utils.GetMsgByCommand("getnewaddress",nil)
//	if err != nil {
//		return nil, err
//	}
//	return Rpcresult.Result, nil
//
//}

//
/*func GetBlockHeader(hash interface{})(interface{},error)  {
	jsonMes, err := utils.PrepareJsonStr("getblockheader", nil)
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
	return Rpcresult.Result, nil*/

