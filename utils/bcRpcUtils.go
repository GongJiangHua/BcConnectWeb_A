package utils

import (
	"BcConnectWeb_A/Constants"
	"BcConnectWeb_A/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)
/**
1、准备一个json数据
2、使用http链接的post请求，发送json数据
3、接受http响应
4、为post请求添加请求头
5、处理返回的结果
 */


//1、准备一个json数据用于向
func PrepareJsonStr(method string,params ...interface{}) ([]byte,error) {
	//json操作：序列化、反序列化
	rpcReq := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Jsonrpc: "2.0",
		Method:  method,
		Params: params,
	}
	//对结构体类型进行序列化
	rpcBytes,err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return rpcBytes,nil
}


func SendRPCPost(jsonMes []byte) (*entity.RPCResult,error) {
	//client：客户端；客户端用于发起请求
	client := http.Client{}//实例化一个客户端

	//实例化一个请求
	request,err := http.NewRequest("POST",Constants.RPCURL,bytes.NewBuffer(jsonMes))
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	//给post请求添加请求头设置信息
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("Authorization","Basic " + Base64Str(Constants.RPCUSER+":"+Constants.RPCPASSWORD))

	//使用客户端发送请求
	response,err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}

	//通过response获取响应的数据
	code := response.StatusCode
	resultBytes,err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	//fmt.Println(resultBytes)

	if code == http.StatusOK {

		//fmt.Println("这",string(resultBytes))
		var result *entity.RPCResult
		err = json.Unmarshal(resultBytes,&result)
		if err != nil {
			fmt.Println(err.Error())
			return nil,nil
		}

		//反序列化正常，没有出现错误
		//fmt.Println("功能调用结果：",result.Result)
		return result,nil
	}else {
		fmt.Println("请求失败，错误状态码是：",code)
		return nil,err
	}
}
/*btc命令调用封装函数 命令 [参数1，参数2 ...]
method： 比特币节点具体命令
parms ：命令对应的具体参数
return：比特币 Result
*/
func GetMsgByCommand(method string, params ...interface{}) (*entity.RPCResult, error) {
	jsonStr ,_:= PrepareJsonStr(method, params...)
	fmt.Println(jsonStr)
	resBytes,err := SendRPCPost(jsonStr)
	if err != nil {
		return nil,err
	}
	return resBytes,nil
}
//解析返回的结果，并且反序列化成结构体
//func ParseResultBytes(resBytes []byte) (*entity.RPCResult,error) {
//	var Rpcresult entity.RPCResult
//	err := json.Unmarshal(resBytes,&Rpcresult)
//	if err != nil {
//		return nil,err
//	}
//	return &Rpcresult,nil
//}