package main

import (
	"encoding/json"
	"fmt"
	_ "strconv"
	_ "errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strings"
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
	fmt.Printf("Func addRecord begin===== \n")

	var param string
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}


	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		// 如果key不存在
		fmt.Println("Key Not Found， Add New Record")
		paramMap := make(map[string]interface{})
		param += args[2] + "_" + args[3]
		paramMap[args[1]] = param
		str, err := json.Marshal(paramMap)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("addRecord Params:%s\n", string(str))

		err1:=APIstub.PutState(args[0], []byte(str))
		if err1 != nil{
			return shim.Error("wirteIn error")
		}
	} else{
		// 如果key存在
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(containerAsBytes), &dat); err == nil {
			fmt.Println(dat)
			paramMap := make(map[string]interface{})
			param += args[2] + "_" + args[3]
			paramMap[args[1]] = param
			str, err := json.Marshal(paramMap)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("addRecord Params:%s\n", string(str))

			err1:=APIstub.PutState(args[0], []byte(str))
			if err1 != nil{
				return shim.Error("wirteIn error")
			}
		} else {
			fmt.Println(err)
		}

	}

	return shim.Success(nil)
}

func (s *SmartContract) getRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	fmt.Printf("Func getRecord begin===== \n")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		return shim.Error("getRecord === Could not locate container")
	}
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(containerAsBytes), &dat); err == nil {
		fmt.Println(dat)
	} else {
		fmt.Println(err)
	}
	var value string
	// 查找键值是否存在
	if v, ok := dat[args[1]]; ok {
		value = v.(string)
		fmt.Println("Key Found\t" + v.(string))
		fmt.Println("Key Found\t" + value)
	} else {
		fmt.Println("Key Not Found")
	}

	res := strings.Split(value, "_")

	fmt.Printf("Query Response:%s\n", res[0])

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