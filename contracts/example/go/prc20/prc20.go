/*
	This file is part of go-palletone.
	go-palletone is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
	go-palletone is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.
	You should have received a copy of the GNU General Public License
	along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
*/
/*
 * Copyright IBM Corp. All Rights Reserved.
 * @author PalletOne core developers <dev@pallet.one>
 * @date 2018
 */

package prc20

import (
	"encoding/json"
	"fmt"
	"strconv"

	//	cm "github.com/palletone/go-palletone/contracts/modules"
	"github.com/palletone/go-palletone/contracts/shim"
	pb "github.com/palletone/go-palletone/core/vmContractPub/protos/peer"
	dm "github.com/palletone/go-palletone/dag/modules"
)

const symbolsKey = "symbols"

type PRC20 struct {
}

type Symbols struct {
	nameAddrs map[string]string `json:"nameAddrs"`
}

func (p *PRC20) Init(stub shim.ChaincodeStubInterface) pb.Response {
	symbols := Symbols{nameAddrs: map[string]string{}}
	err := setSymbols(&symbols, stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (p *PRC20) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	f, args := stub.GetFunctionAndParameters()

	switch f {
	case "createToken":
		return createToken(args, stub)
	default:
		jsonResp := "{\"Error\":\"Unknown function " + f + "\"}"
		return shim.Error(jsonResp)
	}
}

func setSymbols(symbols *Symbols, stub shim.ChaincodeStubInterface) error {
	val, err := json.Marshal(symbols)
	if err != nil {
		return err
	}
	err = stub.PutState(symbolsKey, val)
	return err
}

func getSymbols(stub shim.ChaincodeStubInterface) (*Symbols, error) {
	//
	symbolsBytes, err := stub.GetState(symbolsKey)
	if err != nil {
		return nil, err
	}
	//
	var symbols Symbols
	err = json.Unmarshal(symbolsBytes, &symbols)
	if err != nil {
		return nil, err
	}

	return &symbols, nil
}

func createToken(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	//params check
	if len(args) < 4 {
		return shim.Error("need 4 args (Name,Symbol,Decimals,TotalSupply,[SupplyAddress])")
	}

	//==== convert params to token information
	var fungible dm.FungibleToken
	//name symbol
	fungible.Name = args[0]
	fungible.Symbol = args[1]
	//transfer how to use ?
	fungible.Decimals = []byte(args[2])[0]
	//supply amount
	toalSupply, err := strconv.ParseUint(args[3], 10, 64)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to convert total supply\"}"
		return shim.Error(jsonResp)
	}
	fungible.TotalSupply = toalSupply
	//address of supply
	if len(args) > 4 {
		fungible.SupplyAddress = args[4]
	}

	//check name is only or not
	symbols, err := getSymbols(stub)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get symbols\"}"
		return shim.Error(jsonResp)
	}
	if _, ok := symbols.nameAddrs[fungible.Symbol]; ok {
		jsonResp := "{\"Error\":\"The symbol have been used\"}"
		return shim.Success([]byte(jsonResp))
	}

	//convert to json
	createJson, err := json.Marshal(fungible)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to generate token Json\"}"
		return shim.Error(jsonResp)
	}
	//==== init tokendefine
	var tokenDefine dm.TokenDefine
	tokenDefine.TokenDefineJson = createJson
	tokenDefine.TokenType = 0
	//get invoke address
	createaddr, err := stub.GetInvokeAddress()
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get invoke address\"}"
		return shim.Error(jsonResp)
	}
	tokenDefine.Creator = createaddr

	//return result
	result, err := json.Marshal(tokenDefine)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to Marshal result\"}"
		return shim.Error(jsonResp)
	}

	//set token define
	err = stub.DefineToken(byte(0), result)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to call stub.DefineToken\"}"
		return shim.Error(jsonResp)
	}

	//last put state
	symbols.nameAddrs[fungible.Symbol] = fungible.SupplyAddress
	err = setSymbols(symbols, stub)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to set symbols\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(result) //test
}

func putval(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	if len(args) < 2 {
		return shim.Error("need 2 args (key and a value)")
	}
	key := args[0]
	val := args[1]
	// Get the state from the ledger
	valbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get val for " + key + "\"}"
		return shim.Error(jsonResp)
	}
	fmt.Println("==== valOld demo ==== ", key, string(valbytes))
	// Write the state to the ledger
	err = stub.PutState(key, []byte(val))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func getval(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key to query")
	}
	key := args[0]
	// Get the state from the ledger
	valbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(jsonResp)
	}
	fmt.Println("==== valOld demo ==== ", key, string(valbytes))
	if valbytes == nil {
		jsonResp := "{\"Error\":\"Nil val for " + key + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valbytes)
}