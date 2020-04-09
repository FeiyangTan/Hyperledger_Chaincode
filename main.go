//该链码基于hyperledger fabric 1.4 开发
//实现资产资产在用户之间的共享
package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type AssertsExchangeCC struct{}

//User 用户
type User struct {
	Name   string   `json:"name"`
	ID     string   `json:"id"`
	Assets []string `json:"assets"`
}

//Asset 资产
type Asset struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Infor string `json:"infor"`
}

//AssetExchange 资产转移记录
type AssetExchangeRecord struct {
	//Id 资产id
	ID            string `json:"id"`
	OriginOwnerID string `json:"origin_owner_id"`
	NewOwnerID    string `json:"new_owner_id"`
}

//开始
func main() {
	err := shim.Start(new(AssertsExchangeCC))
	errHandle("Starting AssertsExchange chaincode container error: ", err)
}

//Init 创建Chaincode container
func (a *AssertsExchangeCC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Init success"))
}

//Invoke 操作/查询Chaincode
func (a *AssertsExchangeCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()

	switch funcName {
	case "userRegister":
		return userRegister(stub, args)
	case "userDelete":
		return userDelete(stub, args)
	case "assetRegister":
		return userRegister(stub, args)
	case "assetDelete":
		return assetDelete(stub, args)
	case "assetExchange":
		return assetExchange(stub, args)
	case "queryUser":
		return queryUser(stub, args)
	case "queryAsset":
		return queryAsset(stub, args)
	case "queryAssetExchangeRecord":
		return queryAssetExchangeRecord(stub, args)
	default:
		return shim.Error("unexpect function name: " + funcName)
	}
}

func errHandle(infor string, err error) {
	if err != nil {
		fmt.Println(infor, err)
	}
}
