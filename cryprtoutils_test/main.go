package main

import (
	"encoding/json"
	"fmt"

	"github.com/thomasks/eju-chaincodes/cryptoutils"
)

var testdata = `[{
	"id": 52235,
	"createTime": 1530668028893,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 10,
		"code": "YYRB_YSFJ",
		"name": "夜审房金",
		"level": 2,
		"parentId": 9
	},
	"demension": {
		"id": 1,
		"code": "YYRB_Label_5_0",
		"name": "发生代码"
	},
	"dataPacket": {
		"id": 135,
		"filePath": "http://demo1.essintra.ejucloud.cn/fbcstore/1/20180704_093344_3861.xls",
		"up2Chain": false
	},
	"beginDate": 1517414400000,
	"endDate": 1517500800000,
	"period": "d",
	"value": "1001.0",
	"indicatorIdAndBeginDate": 1517414400010
}]`

var testcd = `[{"level":"GROUP","cryptoFields":["value"]}]`

func main() {

	var cds []cryptoutils.CryptoDescriptor
	if err := json.Unmarshal([]byte(testcd), &cds); err != nil {
		fmt.Println("meet error." + err.Error())
	}

	var rawDataMapArr []map[string]interface{}
	if err := json.Unmarshal([]byte(testdata), &rawDataMapArr); err != nil {
		fmt.Println("unmarshal value error: " + err.Error())
	}

	fmt.Printf("before crypto %#v\n\n", rawDataMapArr)
	//crypto value in each level
	for _, rawDataMap := range rawDataMapArr {
		cryptoutils.CryptoDataByDescriptor(nil, rawDataMap, cds)
		//attris := cd.CryptoFields
		//level := cd.Level
		//getCryptoKey4ChannelLevel(level, stub.GetChannelID)
	}
	fmt.Printf("after crypto %#v\n\n", rawDataMapArr)

	for _, rawDataMap := range rawDataMapArr {
		cryptoutils.DecryptoDataByDescriptor(nil, rawDataMap, cds)
		//attris := cd.CryptoFields
		//level := cd.Level
		//getCryptoKey4ChannelLevel(level, stub.GetChannelID)
	}
	fmt.Printf("after encrpt %#v\n\n", rawDataMapArr)
}
