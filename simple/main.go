package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/thomasks/eju-chaincodes/cryptoutils"
)

// ToChaincodergs comment
func ToChaincodergs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

// Chaincode comment
type Chaincode struct {
}

//{"Args":["attr", "name"]}'
func (t *Chaincode) attr(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("get attr: ", args[0])
	value, ok, err := cid.GetAttributeValue(stub, args[0])
	if err != nil {
		return shim.Error("get attr error: " + err.Error())
	}

	if ok == false {
		value = "not found"
	}
	bytes, err := json.Marshal(value)
	if err != nil {
		return shim.Error("json marshal error: " + err.Error())
	}
	return shim.Success(bytes)
}

//{"Args":["creator"]}'
func (t *Chaincode) creator(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("creator: ", args)
	bytes, err := stub.GetCreator()
	if err != nil {
		return shim.Error("get creator error: " + err.Error())
	}

	var creator msp.SerializedIdentity
	if err := json.Unmarshal(bytes, &creator); err != nil {
		return shim.Error("unmarshal creator error: " + err.Error())
	}
	return shim.Success(nil)
}

//{"Args":["call","chaincode","method"...]}'
func (t *Chaincode) call(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("call: ", args)
	subArgs := args[1:]
	return stub.InvokeChaincode(args[0], ToChaincodergs(subArgs...), stub.GetChannelID())
}

//{"Args":["append","key", ...]}'
func (t *Chaincode) append(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	key := args[0]
	value := args[1:]
	var data []string

	bytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("query " + key + " fail: " + err.Error())
	}

	if bytes != nil {
		if err := json.Unmarshal(bytes, &data); err != nil {
			return shim.Error(err.Error())
		}
	}

	data = append(data, value...)
	newBytes, err := json.Marshal(data)
	if err != nil {
	}

	if err := stub.PutState(key, newBytes); err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//{"Args":["query_chaincode","chaincode","key"]}'
func (t *Chaincode) queryChaincode(stub shim.ChaincodeStubInterface, chaincode, key string) pb.Response {
	fmt.Printf("query %s in %s\n", key, chaincode)
	return stub.InvokeChaincode(chaincode, ToChaincodergs("query", key), stub.GetChannelID())
}

//{"Args":["write_chaincode","chaincode","key","value"]}'
func (t *Chaincode) writeChaincode(stub shim.ChaincodeStubInterface, chaincode, key, value string) pb.Response {
	fmt.Printf("write %s to %s, value is %s\n", key, chaincode, value)
	return stub.InvokeChaincode(chaincode, ToChaincodergs("write", key, value), stub.GetChannelID())
}

//{"Args":["query","key"]}'
func (t *Chaincode) query(stub shim.ChaincodeStubInterface, key string) pb.Response {
	fmt.Printf("query %s\n", key)
	bytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("query fail " + err.Error())
	}
	return shim.Success(bytes)
}

//{"Args":["history","key"]}'
func (t *Chaincode) history(stub shim.ChaincodeStubInterface, key string) pb.Response {
	fmt.Printf("history %s\n", key)
	iter, err := stub.GetHistoryForKey(key)
	defer iter.Close()
	if err != nil {
		return shim.Error("query fail " + err.Error())
	}

	values := make(map[string]string)

	for iter.HasNext() {
		fmt.Printf("next\n")
		if kv, err := iter.Next(); err == nil {
			fmt.Printf("id: %s value: %s\n", kv.TxId, kv.Value)
			values[kv.TxId] = string(kv.Value)
		}
		if err != nil {
			return shim.Error("iterator history fail: " + err.Error())
		}
	}

	bytes, err := json.Marshal(values)
	if err != nil {
		return shim.Error("json marshal fail: " + err.Error())
	}

	return shim.Success(bytes)
}

//{"Args":["write","key","value"]}'
func (t *Chaincode) write(stub shim.ChaincodeStubInterface, key, value string) pb.Response {
	fmt.Printf("write %s, value is %s\n", key, value)
	if err := stub.PutState(key, []byte(value)); err != nil {
		return shim.Error("write fail " + err.Error())
	}
	return shim.Success(nil)
}

//{"Args":["writeMultiSegData","key","value",SegDescriptor]}
func (t *Chaincode) writeMultiSegData(stub shim.ChaincodeStubInterface, key, value, cryptoDescriptor string) pb.Response {
	fmt.Printf("write %s,value is %s,SegDescriptor is %s\n", key, value, cryptoDescriptor)

	var cds []cryptoutils.CryptoDescriptor
	if err := json.Unmarshal([]byte(cryptoDescriptor), &cds); err != nil {
		return shim.Error("unmarshal cryptoDescriptor error: " + err.Error())
	}

	var rawDataMapArr []map[string]interface{}
	if err := json.Unmarshal([]byte(value), &rawDataMapArr); err != nil {
		return shim.Error("unmarshal value error: " + err.Error())
	}

	//crypto value in each level
	for _, rawDataMap := range rawDataMapArr {
		cryptoutils.CryptoDataByDescriptor(stub, rawDataMap, cds)
		//attris := cd.CryptoFields
		//level := cd.Level
		//getCryptoKey4ChannelLevel(level, stub.GetChannelID)
	}

	blockHead := BlockHead{
		CryptoDescriptor: cryptoDescriptor,
	}

	writeTo := HeadBodyBlock{
		Head: blockHead,
		Body: rawDataMapArr,
	}

	bytes, err := json.Marshal(writeTo)
	if err != nil {
		return shim.Error("json marshal error: " + err.Error())
	}
	if err := stub.PutState(key, bytes); err != nil {
		return shim.Error("write fail " + err.Error())
	}
	return shim.Success(nil)
}

//{"Args":["readMultiSegData","key","value",SegDescriptor]}
func (t *Chaincode) readMultiSegData(stub shim.ChaincodeStubInterface, key string) pb.Response {
	fmt.Printf("key %s\n", key)
	bytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("query fail " + err.Error())
	}

	var readTo = new(HeadBodyBlock)
	if err := json.Unmarshal(bytes, readTo); err != nil {
		return shim.Error("unmarshal rawBytes error: " + err.Error())
	}
	var cds []cryptoutils.CryptoDescriptor
	if err := json.Unmarshal([]byte(readTo.Head.CryptoDescriptor), &cds); err != nil {
		return shim.Error("unmarshal cryptoDescriptor error: " + err.Error())
	}

	//crypto value in each level
	for _, rawDataMap := range readTo.Body {
		cryptoutils.DecryptoDataByDescriptor(stub, rawDataMap, cds)
		//attris := cd.CryptoFields
		//level := cd.Level
		//getCryptoKey4ChannelLevel(level, stub.GetChannelID)
	}

	ret, err2 := json.Marshal(readTo.Body)
	if err2 != nil {
		return shim.Error("json marshal error: " + err.Error())
	}
	return shim.Success(ret)
}

//BlockHead descr
type BlockHead struct {
	CryptoDescriptor string `json:"cryptoDescriptor"`
}

//HeadBodyBlock desc
type HeadBodyBlock struct {
	Head BlockHead                `json:"head"`
	Body []map[string]interface{} `json:"body"`
}

//Init {"Args":["init"]}
func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init Chaincode Chaincode")
	return shim.Success(nil)
}

//Invoke {"write":["key","value"]}
//Invoke {"writeMultiSegData":["key","value","SegDataDescriptor"]}
//
func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "creator":
		return t.creator(stub, args)
	case "call": //调用
		return t.call(stub, args)
	case "append": //追加
		return t.append(stub, args)
	case "attr":
		if len(args) != 1 {
			return shim.Error("parametes's number is wrong")
		}
		return t.attr(stub, args)
	case "query": //查询
		if len(args) != 1 {
			return shim.Error("parametes's number is wrong")
		}
		return t.query(stub, args[0])
	case "history": //查询
		if len(args) != 1 {
			return shim.Error("parametes's number is wrong")
		}
		return t.history(stub, args[0])
	case "write": //写入
		if len(args) != 2 {
			return shim.Error("parametes's number is wrong")
		}
		return t.write(stub, args[0], args[1])
	case "writeMultiSegData": //写入
		if len(args) != 3 {
			return shim.Error("parametes's number is wrong")
		}
		return t.writeMultiSegData(stub, args[0], args[1], args[2])
	case "readMultiSegData": //写入
		if len(args) != 1 {
			return shim.Error("parametes's number is wrong")
		}
		return t.readMultiSegData(stub, args[0])
	case "query_chaincode":
		if len(args) != 2 {
			return shim.Error("parametes's number is wrong")
		}
		return t.queryChaincode(stub, args[0], args[1])
	case "write_chaincode":
		if len(args) != 3 {
			return shim.Error("parametes's number is wrong")
		}
		return t.writeChaincode(stub, args[0], args[1], args[2])
	default:
		return shim.Error("Invalid invoke function name.")
	}
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting Chaincode chaincode: %s", err)
	}
}
