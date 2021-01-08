package btcService

import (
	"BcConnectWeb_A/entity"
	"BcConnectWeb_A/utils"
	"github.com/mitchellh/mapstructure"
)

//该包用来封装bitcoin命令

//用于查询区块链中的区块数量
func GetBlockCount() (interface{}, error) {
	Rpcresult,err := utils.GetMsgByCommand("getblockcount")
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil
}

//用于查询区块链中最新区块的hash

func GetBestBlockHash() (interface{}, error) {
	Rpcresult, err := utils.GetMsgByCommand("getbestblockhash")
	if err != nil {
		return nil, err
	}
	return Rpcresult.Result, nil
}

//根据区块高度查询区块hash
func GetBlockHashByHeight(height int) (interface{}, error) {
	Rpcresult,err := utils.GetMsgByCommand("getblockhash",height)
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil
}

//用于获取当前区块的难度
func GetDifficulty() (interface{}, error) {
	Rpcresult, err := utils.GetMsgByCommand("getdifficulty")
	if err != nil {
		return nil, err
	}
	return Rpcresult.Result, nil
}

//用于获取当前区块链信息
func GetBlockChainInfo() (*entity.BlockChainInfo,error) {
	Rpcresult,err := utils.GetMsgByCommand("getblockchaininfo")
	if err != nil {
		return nil,err
	}
	var blockchaininfo entity.BlockChainInfo
	err = mapstructure.Decode(Rpcresult.Result,&blockchaininfo)
	if err != nil {
		return nil,err
	}
	return &blockchaininfo,nil
}

////生成一个新地址
func GetNewAddress()(interface{},error)  {
	Rpcresult, err := utils.GetMsgByCommand("getnewaddress")
	if err != nil {
		return nil, err
	}
	return Rpcresult.Result, nil

}

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

//用于获取钱包信息
func GetWalletInfo() (*entity.WalletInfo,error) {
	Rpcresult,err := utils.GetMsgByCommand("getwalletinfo")
	if err != nil {
		return nil,err
	}
	var blockchaininfo entity.WalletInfo
	err = mapstructure.Decode(Rpcresult.Result,&blockchaininfo)
	if err != nil {
		return nil,err
	}
	return &blockchaininfo,nil
}