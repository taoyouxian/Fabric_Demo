# Fabric_Demo
Coding happily.

# Intro 
本团队是将`container-app`拷贝到`src/github.com/hyperledger/fabric-samples/chaincode/`目录下，然后实现的测试。

# Run
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

# Contact us
For feedback and questions, feel free to email us:
- Youxian Tao taoyouxian@aliyun.com
- Longying Wu navicate@163.com

Welcome to contribute and submit pull requests :)

Our repository([Fabric_Demo](https://github.com/taoyouxian/Fabric_Demo.git)) is private currently. 

# Reference
[相关参考](https://segmentfault.com/a/1190000013080245)