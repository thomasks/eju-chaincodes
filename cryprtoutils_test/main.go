package main

import (
	"encoding/json"
	"fmt"
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
	var rawDataMapArr []map[string]interface{}
	if err := json.Unmarshal([]byte(testdata), &rawDataMapArr); err != nil {
		fmt.Println("unmarshal value error: " + err.Error())
	}
	for _, rawDataMap := range rawDataMapArr {
		id := rawDataMap["id"]

		if id != nil {
			idStr := fmt.Sprintf("%s-%v", "1_20170107_1#1879", id)
			fmt.Printf("rawDataMap id type is [%T]\n rawDataMap id value is [%v]\n", id, id)
			fmt.Printf("rawDataMap idStr type is [%T]\n rawDataMap idStr value is [%v]\n", idStr, idStr)
		}

	}
}

/*func testFlow() {
	var cds []cryptoutils.CryptoDescriptor
	if err := json.Unmarshal([]byte(testcd), &cds); err != nil {
		fmt.Println("meet error." + err.Error())
	}

	var rawDataMapArr []map[string]interface{}
	if err := json.Unmarshal([]byte(testdata), &rawDataMapArr); err != nil {
		fmt.Println("unmarshal value error: " + err.Error())
	}

	fmt.Printf("before crypto %#v\n\n", rawDataMapArr)
	for _, rawDataMap := range rawDataMapArr {
		cryptoutils.CryptoDataByDescriptor(nil, rawDataMap, cds)
	}
	fmt.Printf("after crypto %#v\n\n", rawDataMapArr)

	for _, rawDataMap := range rawDataMapArr {
		cryptoutils.DecryptoDataByDescriptor(nil, rawDataMap, cds)
	}
	fmt.Printf("after encrpt %#v\n\n", rawDataMapArr)
}*/
