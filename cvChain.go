package main

import (
	"encoding/json"
	"log"
	_ "strconv"
	_ "errors"
	"strings"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
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
		return s.getRecord(APIstub, args)
	} else if function == "encRecord" {
		return s.encRecord(APIstub, args, key[ENCKEY], key[IV])
	} else if function == "decRecord" {
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
	ID := args[0]
	log.Printf("- start getHistoryForKey: %s", ID)

	resultsIterator, err := APIstub.GetHistoryForKey(ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if response == nil {
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

			err1 := APIstub.PutState(args[0], []byte(str))
			if err1 != nil {
				return shim.Error("wirteIn error")
			}

			break
		} else {
			var containerAsBytes string
			containerAsBytes = string(response.Value)
			log.Println("resultsIterator \t" + containerAsBytes)
			// If key found
			log.Println("Key Found, Get Old Value")
			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(containerAsBytes), &dat); err == nil {
				log.Println("Old Value: ")
				log.Print(dat)
				log.Println()
				var value string
				if v, ok := dat[args[1]]; ok {
					value = v.(string)
					log.Println("Key Found\t" + value)
					return shim.Error("wirteIn same year error\t" + value)
				} else {
					log.Println("Key Not Found")
				}

				break
			} else {
				log.Println(err)
			}

		}
	}

	dat := make(map[string]interface{})
	param = args[2] + "_" + args[3]
	dat[args[1]] = param
	str, err := json.Marshal(dat)
	if err != nil {
		log.Println(err)
	}
	log.Printf("addRecord Params:%s\n", string(str))

	err1 := APIstub.PutState(args[0], []byte(str))
	if err1 != nil {
		return shim.Error("wirteIn error")
	}
	return shim.Success(nil)
}

func (s *SmartContract) getRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	log.Printf("Func getRecord begin===== \n")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	resultsIterator, err := APIstub.GetHistoryForKey(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var value string
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var containerAsBytes string
		containerAsBytes = string(response.Value)
		log.Println("resultsIterator \t" + containerAsBytes)

		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(containerAsBytes), &dat); err == nil {
			log.Println(dat)
		} else {
			log.Println(err)
		}
		// Judge key is found or not
		if v, ok := dat[args[1]]; ok {
			value = v.(string)
			log.Println("getRecord Key Found\t" + value)
			log.Printf("getRecord Query Response:%s\n", value)
			break
		} else {
			log.Println("Key Not Found")
		}
	}
	res := strings.Split(value, "_")
	return shim.Success([]byte(res[0]))

}

func (s *SmartContract) encRecord(APIstub shim.ChaincodeStubInterface, args []string, encKey, IV []byte) sc.Response {
	log.Printf("Func encRecord begin===== \n")

	var param string
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	ent, err := entities.NewAES256EncrypterEntity("ID", s.bccspInst, encKey, IV)
	if err != nil {
		return shim.Error(fmt.Sprintf("entities.NewAES256EncrypterEntity failed, err %s", err))
	}

	ID := args[0]
	log.Printf("- start getHistoryForKey: %s", ID)

	resultsIterator, err := APIstub.GetHistoryForKey(ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if response == nil {
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

			break
		} else {
			var containerAsBytes string
			containerAsBytes = string(response.Value)
			log.Println("resultsIterator \t" + containerAsBytes)
			cleartextValue, err := ent.Decrypt(response.Value)
			if err != nil {
				return shim.Error(fmt.Sprintf("getStateAndDecrypt failed, err %+v", err))
			}
			log.Println("resultsIterator \t" + string(cleartextValue))
			// If key found
			log.Println("Key Found, Get Old Value")
			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(cleartextValue), &dat); err == nil {
				log.Println("Old Value: ")
				log.Print(dat)
				log.Println()
				var value string
				if v, ok := dat[args[1]]; ok {
					value = v.(string)
					log.Println("Key Found\t" + value)
					return shim.Error("wirteIn same year error\t" + value)
				} else {
					log.Println("Key Not Found")
				}

				break
			} else {
				log.Println(err)
			}

		}
	}

	dat := make(map[string]interface{})
	param = args[2] + "_" + args[3]
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

	return shim.Success(nil)
}

func (s *SmartContract) decRecord(APIstub shim.ChaincodeStubInterface, args []string, decKey, IV []byte) sc.Response {
	log.Println("Func decRecord begin===== \n")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	ent, err := entities.NewAES256EncrypterEntity("ID", s.bccspInst, decKey, IV)

	ID := args[0]

	resultsIterator, err := APIstub.GetHistoryForKey(ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var value string
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		var containerAsBytes string
		containerAsBytes = string(response.Value)
		log.Println("resultsIterator \t" + containerAsBytes)

		cleartextValue, err := ent.Decrypt(response.Value)
		if err != nil {
			return shim.Error(fmt.Sprintf("getStateAndDecrypt failed, err %+v", err))
		}
		log.Println(cleartextValue)

		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(cleartextValue), &dat); err == nil {
			log.Println(dat)
		} else {
			log.Println(err)
		}
		// Judge key is found or not
		if v, ok := dat[args[1]]; ok {
			value = v.(string)
			log.Println("decRecord Key Found\t" + value)
			log.Printf("Query Response:%s\n", value)
			break
		} else {
			log.Println("decRecord Key Not Found")
		}

	}
	res := strings.Split(value, "_")

	return shim.Success([]byte(res[0]))
}

func main() {
	factory.InitFactories(nil)
	err := shim.Start(&SmartContract{factory.GetDefault()})

	//err := shim.Start(new(SmartContract))
	if err != nil {
		log.Printf("Error creating new Smart Contract: %s", err)
	}
}