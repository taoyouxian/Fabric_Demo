package main

import (
	"encoding/json"
	"log"
	_ "strconv"
	_ "errors"
	"strings"
	"github.com/hyperledger/fabric/bccsp"
	_ "github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
	sc "github.com/hyperledger/fabric/protos/peer"
	"fmt"
)
const DECKEY = "DECKEY"
const VERKEY = "VERKEY"
const ENCKEY = "ENCKEY"
const SIGKEY = "SIGKEY"
const IV = "IV"

type SmartContract struct {
	bccspInst bccsp.BCCSP
}

type Container struct {
	X string `json:"x"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	key, _ := APIstub.GetTransient()

	if function == "addRecord" {
		return s.addRecord(APIstub, args)
	} else if function == "getRecord" {
		return s.getRecord(APIstub,args)
	} else if function == "encRecord" {
		for key, value := range key {
			log.Println("Key:", key, "Value:", value)
		}

		return s.encRecord(APIstub, args, key[ENCKEY], key[IV])
	} else if function == "decRecord" {
		for key, value := range key {
			log.Println("Key:", key, "Value:", value)
		}

		return s.decRecord(APIstub, args, key[DECKEY], key[IV])
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) addRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	log.Printf("Func addRecord begin===== \n")

	var param string
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}


	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		// If key not found
		log.Println("Key Not Found, Add New Record")
		paramMap := make(map[string]interface{})
		param += args[2] + "_" + args[3]
		paramMap[args[1]] = param
		str, err := json.Marshal(paramMap)
		if err != nil {
			log.Println(err)
		}
		log.Printf("addRecord Params:%s\n", string(str))

		err1:=APIstub.PutState(args[0], []byte(str))
		if err1 != nil{
			return shim.Error("wirteIn error")
		}
	} else{
		// If key found
		log.Println("Key Found, Get Old Value")
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(containerAsBytes), &dat); err == nil {
			log.Println("Old Value: ")
			log.Print(dat)
			log.Println()
			param += args[2] + "_" + args[3]
			dat[args[1]] = param
			str, err := json.Marshal(dat)
			if err != nil {
				log.Println(err)
			}
			log.Printf("addRecord Params:%s\n", string(str))

			err1:=APIstub.PutState(args[0], []byte(str))
			if err1 != nil{
				return shim.Error("wirteIn error")
			}
		} else {
			log.Println(err)
		}

	}

	return shim.Success(nil)
}

func (s *SmartContract) getRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	log.Printf("Func getRecord begin===== \n")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		return shim.Error("getRecord === Could not locate container")
	}
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(containerAsBytes), &dat); err == nil {
		log.Println(dat)
	} else {
		log.Println(err)
	}
	var value string
	// Judge key is found or not
	if v, ok := dat[args[1]]; ok {
		value = v.(string)
		log.Println("Key Found\t" + value)
	} else {
		log.Println("Key Not Found")
	}

	res := strings.Split(value, "_")

	log.Printf("Query Response:%s\n", value)

	return shim.Success([]byte(res[0]))
}


func (s *SmartContract) encRecord(APIstub shim.ChaincodeStubInterface, args []string, encKey, IV []byte) sc.Response {
	log.Printf("Func encRecord begin===== \n")

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	ent, err := entities.NewAES256EncrypterEntity("ID", s.bccspInst, encKey, IV)
	if err != nil {
		return shim.Error(fmt.Sprintf("entities.NewAES256EncrypterEntity failed, err %s", err))
	}
	var param string
	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		// If key not found
		log.Println("Key Not Found, Add New Record")
		paramMap := make(map[string]interface{})
		param += args[2] + "_" + args[3]
		paramMap[args[1]] = param
		str, err := json.Marshal(paramMap)
		if err != nil {
			log.Println(err)
		}
		log.Printf("addRecord Params:%s\n", string(str))

		// here, we encrypt Value and assign it to key
		err = encryptAndPutState(APIstub, ent, args[0], []byte(str))
		if err != nil {
			return shim.Error(fmt.Sprintf("encryptAndPutState failed, err %+v", err))
		}

	} else {
		// If key found
		log.Println("Key Found, Get Old Value")
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(containerAsBytes), &dat); err == nil {
			log.Println("Old Value: ")
			log.Print(dat)
			log.Println()
			param += args[2] + "_" + args[3]
			dat[args[1]] = param
			str, err := json.Marshal(dat)
			if err != nil {
				log.Println(err)
			}
			log.Printf("addRecord Params:%s\n", string(str))

			err = encryptAndPutState(APIstub, ent, args[0], []byte(str))
			if err != nil {
				return shim.Error(fmt.Sprintf("encryptAndPutState failed, err %+v", err))
			}
		} else {
			log.Println(err)
		}

	}

	return shim.Success(nil)
}

func (s *SmartContract) decRecord(APIstub shim.ChaincodeStubInterface, args []string, decKey, IV []byte) sc.Response {
	log.Printf("Func decRecord begin===== \n")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	ent, err := entities.NewAES256EncrypterEntity("ID", s.bccspInst, decKey, IV)

	key := args[0]

	// here we decrypt the state associated to key
	cleartextValue, err := getStateAndDecrypt(APIstub, ent, key)
	// 这里需要循环判断输出与输入年份一致的信息
	if err != nil {
		return shim.Error(fmt.Sprintf("getStateAndDecrypt failed, err %+v", err))
	}
	log.Println(cleartextValue)

	// here we return the decrypted value as a result
	return shim.Success(cleartextValue)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		log.Printf("Error creating new Smart Contract: %s", err)
	}
}