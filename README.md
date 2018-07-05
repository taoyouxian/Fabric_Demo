# Fabric_Demo
Coding happily.

## Intro 
`人工链圈` - **Artificial chain**

## Run release-1.1
### Start the network(In Terminal)
```
/home/tao/software/opt/Go/src/github.com/hyperledger/fabric-samples/
docker-compose -f docker-compose-cli.yaml up -d
```
- this copys the chaincode dir
```
cd /home/tao/software/opt/Go/src/
docker cp cvChain 49f92f126b1d:/opt/gopath/src/github.com/
```

### Start the network(In Docker)
- this sets configuration
```
docker exec -it cli bash
export CHANNEL_NAME=mychannel
export NAME=cc
export ENCKEY=`openssl rand 32 -base64` && DECKEY=$ENCKEY
export IV=`openssl rand 16 -base64`
```
- Create & Join Channel
```
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```
- Install & Instantiate Chaincode
```
peer chaincode install -n $NAME -v 1.0 -p github.com/cvChain/
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n $NAME -v 1.0 -c '{"Args":["init",""]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')"
```
- Invoke & Query

**`addRecord`**
```
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["addRecord","2018","2018","college1","bachelor"]}'
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["addRecord","2018","2019","college2","bachelor2"]}'
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["addRecord","2018","2018","college3","bachelor3"]}'
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["addRecord","2017","2018","ruc","ruc_bachelor2"]}'
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["addRecord","1","2","ruc12","ruc12_bachelor2"]}'
```
**`getRecord`**
```
peer chaincode query -C $CHANNEL_NAME -n $NAME -c '{"Args":["getRecord","2018", "2019"]}'
peer chaincode query -C $CHANNEL_NAME -n $NAME -c '{"Args":["getRecord","2018", "2020"]}'
peer chaincode query -C $CHANNEL_NAME -n $NAME -c '{"Args":["getRecord","2017", "2018"]}'
```
**`encRecord`**
```
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["encRecord","2017","2017","college1","bachelor"]}' --transient "{\"ENCKEY\":\"$ENCKEY\",\"IV\":\"$IV\"}"
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["encRecord","2015","2016","college1","bachelor"]}' --transient "{\"ENCKEY\":\"$ENCKEY\",\"IV\":\"$IV\"}"
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["encRecord","2017","2018","college2","bachelor2"]}' --transient "{\"ENCKEY\":\"$ENCKEY\",\"IV\":\"$IV\"}"
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["encRecord","2017","2017","college1","bachelor"]}' --transient "{\"ENCKEY\":\"$ENCKEY\",\"IV\":\"$IV\"}"
peer chaincode invoke -o orderer.example.com:7050  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  -C $CHANNEL_NAME -n $NAME -c '{"Args":["encRecord","1","1","ruc1","ruc1bachelor"]}' --transient "{\"ENCKEY\":\"$ENCKEY\",\"IV\":\"$IV\"}"
```
**`decRecord`**
```
peer chaincode query -C $CHANNEL_NAME -n $NAME -c '{"Args":["decRecord", "2017", "2017"]}' --transient "{\"DECKEY\":\"$DECKEY\"}"
peer chaincode query -C $CHANNEL_NAME -n $NAME -c '{"Args":["decRecord", "2017", "2018"]}' --transient "{\"DECKEY\":\"$DECKEY\"}"
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
2. `Building Your First Network`失败操作
```
./byfn.sh -m down
docker network prune
service docker  restart
./byfn.sh -m generate
./byfn.sh -m up
```

## Reference


