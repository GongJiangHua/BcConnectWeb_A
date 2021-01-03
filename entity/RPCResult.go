package entity

import (
	"BitCoinConnect/entity"
	"encoding/json"
)

type RPCResult struct {
	Id int64 `json:"id"`
	Error string `json:"error"`
	Result interface{} `json:"result"`
}
func ParseResultBytes(resBytes []byte) (*entity.RPCResult,error) {
	Rpcresult := entity.RPCResult{}
	err := json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return &Rpcresult,nil
}