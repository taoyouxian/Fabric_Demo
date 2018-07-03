# Fabric_Demo
Coding happily.

## Intro 
`人工链圈`

## Run release-1.1
docker-compose -f docker-compose-cli.yaml up

docker cp Fabric_Demo ContainerID:/opt/gopath/src/github.com/

docker exec -it cli bash

- this installs the Go chaincode
    - self-defined
```
peer chaincode install -n tc -v 1.0 -p github.com/Fabric_Demo/chaincode/
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n tc -v 1.0 -c '{"Args":["init",""]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')"

peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n tc -c '{"Args":["addRecord","1001","1999","college1","bachelor"]}'
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n tc -c '{"Args":["addRecord","1001", "2003","institute1","master"]}'
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n tc -c '{"Args":["addRecord","1001", "2006","corp1", "engineer"]}'
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n tc -c '{"Args":["addRecord","1002", "2017","ruc", "teacher"]}'

peer chaincode query -C $CHANNEL_NAME -n tc -c '{"Args":["getRecord","1001", "2006"]}'
peer chaincode query -C $CHANNEL_NAME -n tc -c '{"Args":["getRecord","1001", "2003"]}'
peer chaincode query -C $CHANNEL_NAME -n tc -c '{"Args":["getRecord","1002", "2017"]}'
```

    - official
```
export CHANNEL_NAME=mychannel

peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')"
peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n mycc -c '{"Args":["invoke","a","b","10"]}'
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'

peer chaincode install -n myc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n myc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')"
```

## Run release-1.0
```
cp chaincode/container-app/container-chaincode.go /home/tao/software/opt/Go/src/github.com/hyperledger/fabric-samples/chaincode/container-app/
```
```
Terminal 1> cd src/github.com/hyperledger/fabric-samples/chaincode-docker-devmode
            docker-compose -f docker-compose-simple.yaml down
			docker-compose -f docker-compose-simple.yaml up
```
```
Terminal 2> docker exec -it chaincode bash
			cd chaincode/container-app/
			rm -rf container-app
			go build
```
```
Terminal 3> docker exec -it cli bash
			cd chaincode/container-app/
			CORE_PEER_ADDRESS=peer:7051 CORE_CHAINCODE_ID_NAME=mycc:0 ./container-app
```
```
Terminal 4> docker exec -it cli bash
			cd /opt/gopath/src
			peer chaincode install -p chaincodedev/chaincode/container-app -n mycc -v 0
			peer chaincode instantiate -n mycc -v 0 -c '{"Args":["init","a","100","b","200"]}' -C myc
			
			peer chaincode invoke -n mycc -c '{"Args":["addRecord","1001","1999","college1","bachelor"]}' -C myc
			peer chaincode invoke -n mycc -c '{"Args":["addRecord","1001", "2003","institute1","master"]}' -C myc
			peer chaincode invoke -n mycc -c '{"Args":["addRecord","1001", "2006","corp1", "engineer"]}' -C myc
			peer chaincode invoke -n mycc -c '{"Args":["addRecord","1002", "2017","ruc", "teacher"]}' -C myc
			
			peer chaincode invoke -n mycc -c '{"Args":["getRecord","1001", "2003"]}' -C myc
			peer chaincode invoke -n mycc -c '{"Args":["getRecord","1002", "2017"]}' -C myc
```

## Contact us
For feedback and questions, feel free to email us:
- Youxian Tao taoyouxian@aliyun.com
- Longying Wu navicate@163.com

Welcome to contribute and submit pull requests :)

Our repository([Fabric_Demo](https://github.com/taoyouxian/Fabric_Demo.git)) is private currently. 

## Q & A
1. Fabric版本弄错，搭建了1.0的环境

## Reference


