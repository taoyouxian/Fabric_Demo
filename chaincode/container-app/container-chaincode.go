package main

import (
	"encoding/json"
	"fmt"
	_ "strconv"
	_ "errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Container struct {
	X string `json:"x"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()

	if function == "addRecord" {
		return s.addRecord(APIstub, args)
	} else if function == "getRecord" {
		return s.getRecord(APIstub,args)
	} else if function == "encRecord" {
		return s.encRecord(APIstub,args)
	} else if function == "decRecord" {
		return s.decRecord(APIstub,args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) addRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	fmt.Printf("Func addRecord: \n")

	var param string
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	param += args[1] + "_" + args[2] + "_" + args[3]
	fmt.Printf("addRecord Params:%s\n", param)

	err:=APIstub.PutState(args[0], []byte(param))
	if err != nil{
		return shim.Error("wirteIn error")
	}

	return shim.Success(nil)
}

func (s *SmartContract) getRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		return shim.Error("Could not locate container")
	}
	fmt.Printf("Query Response:%s\n", containerAsBytes)

	return shim.Success(containerAsBytes)
}

func (s *SmartContract) encRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	containerAsBytes, _ := json.Marshal(args[1])

	err:=APIstub.PutState(args[0],containerAsBytes)
	if err != nil{
		return shim.Error("wirteIn error")
	}

	return shim.Success(nil)
}

func (s *SmartContract) decRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		return shim.Error("Could not locate container")
	}
	return shim.Success(containerAsBytes)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}