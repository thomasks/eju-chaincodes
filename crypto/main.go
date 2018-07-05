package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var itemString = `[{
		"id": 2366,
		"createTime": 1530354842273,
		"merchant": {
			"id": 1,
			"name": "上海源意酒店有限公司"
		},
		"indicator": {
			"id": 21,
			"code": "SYFX_SR_KFSR",
			"name": "客房收入",
			"level": 2,
			"parentId": 34
		},
		"demension": {
			"id": 22,
			"code": "SYFX_JE",
			"name": "金额"
		},
		"dataPacket": {
			"id": 90,
			"filePath": null,
			"up2Chain": false
		},
		"beginDate": 1529251200000,
		"endDate": 1529337600000,
		"period": "d",
		"value": "39673.11",
		"indicatorIdAndBeginDate": 1529251200021
	}]`
	var rawDataMapArr []map[string]interface{}
	err := json.Unmarshal([]byte(itemString), &rawDataMapArr)
	if err == nil {
		fmt.Printf("unmarshal value success.")
		for i := 0; i < len(rawDataMapArr); i++ {
			rawDataMap := rawDataMapArr[i]
			fmt.Printf("value is :%v", rawDataMap["value"])
		}
	} else {
		fmt.Printf("unmarshal value error.%v", err)
	}
}

var jsonStr = `[{
	"id": 2366,
	"createTime": 1530354842273,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 21,
		"code": "SYFX_SR_KFSR",
		"name": "客房收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "39673.11",
	"indicatorIdAndBeginDate": 1529251200021
}, {
	"id": 2367,
	"createTime": 1530354842276,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 21,
		"code": "SYFX_SR_KFSR",
		"name": "客房收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.978",
	"indicatorIdAndBeginDate": 1529251200021
}, {
	"id": 2368,
	"createTime": 1530354842279,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 21,
		"code": "SYFX_SR_KFSR",
		"name": "客房收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200021
}, {
	"id": 2369,
	"createTime": 1530354842282,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 22,
		"code": "SYFX_SR_CYSR",
		"name": "餐饮收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "220.75",
	"indicatorIdAndBeginDate": 1529251200022
}, {
	"id": 2370,
	"createTime": 1530354842285,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 22,
		"code": "SYFX_SR_CYSR",
		"name": "餐饮收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.005",
	"indicatorIdAndBeginDate": 1529251200022
}, {
	"id": 2371,
	"createTime": 1530354842288,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 22,
		"code": "SYFX_SR_CYSR",
		"name": "餐饮收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200022
}, {
	"id": 2372,
	"createTime": 1530354842291,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 23,
		"code": "SYFX_SR_JBJKSR",
		"name": "家宾金卡收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.00",
	"indicatorIdAndBeginDate": 1529251200023
}, {
	"id": 2373,
	"createTime": 1530354842293,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 23,
		"code": "SYFX_SR_JBJKSR",
		"name": "家宾金卡收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.000",
	"indicatorIdAndBeginDate": 1529251200023
}, {
	"id": 2374,
	"createTime": 1530354842296,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 23,
		"code": "SYFX_SR_JBJKSR",
		"name": "家宾金卡收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200023
}, {
	"id": 2375,
	"createTime": 1530354842298,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 24,
		"code": "SYFX_SR_JBPKSR",
		"name": "家宾普卡收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.00",
	"indicatorIdAndBeginDate": 1529251200024
}, {
	"id": 2376,
	"createTime": 1530354842302,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 24,
		"code": "SYFX_SR_JBPKSR",
		"name": "家宾普卡收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.000",
	"indicatorIdAndBeginDate": 1529251200024
}, {
	"id": 2377,
	"createTime": 1530354842306,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 24,
		"code": "SYFX_SR_JBPKSR",
		"name": "家宾普卡收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200024
}, {
	"id": 2378,
	"createTime": 1530354842311,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 25,
		"code": "SYFX_SR_SPSR",
		"name": "商品收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "6.03",
	"indicatorIdAndBeginDate": 1529251200025
}, {
	"id": 2379,
	"createTime": 1530354842315,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 25,
		"code": "SYFX_SR_SPSR",
		"name": "商品收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.000",
	"indicatorIdAndBeginDate": 1529251200025
}, {
	"id": 2380,
	"createTime": 1530354842319,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 25,
		"code": "SYFX_SR_SPSR",
		"name": "商品收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200025
}, {
	"id": 2381,
	"createTime": 1530354842324,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 26,
		"code": "SYFX_SR_DSSR",
		"name": "电商收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "669.83",
	"indicatorIdAndBeginDate": 1529251200026
}, {
	"id": 2382,
	"createTime": 1530354842328,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 26,
		"code": "SYFX_SR_DSSR",
		"name": "电商收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.017",
	"indicatorIdAndBeginDate": 1529251200026
}, {
	"id": 2383,
	"createTime": 1530354842339,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 26,
		"code": "SYFX_SR_DSSR",
		"name": "电商收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200026
}, {
	"id": 2384,
	"createTime": 1530354842344,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 27,
		"code": "SYFX_SR_QTSR",
		"name": "其他收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.00",
	"indicatorIdAndBeginDate": 1529251200027
}, {
	"id": 2385,
	"createTime": 1530354842348,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 27,
		"code": "SYFX_SR_QTSR",
		"name": "其他收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.000",
	"indicatorIdAndBeginDate": 1529251200027
}, {
	"id": 2386,
	"createTime": 1530354842354,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 27,
		"code": "SYFX_SR_QTSR",
		"name": "其他收入",
		"level": 2,
		"parentId": 34
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200027
}, {
	"id": 2387,
	"createTime": 1530354842370,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 30,
		"code": "SYFX_ZC_YYCBZC",
		"name": "营业成本支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "1117.04",
	"indicatorIdAndBeginDate": 1529251200030
}, {
	"id": 2388,
	"createTime": 1530354842375,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 30,
		"code": "SYFX_ZC_YYCBZC",
		"name": "营业成本支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.032",
	"indicatorIdAndBeginDate": 1529251200030
}, {
	"id": 2389,
	"createTime": 1530354842379,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 30,
		"code": "SYFX_ZC_YYCBZC",
		"name": "营业成本支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200030
}, {
	"id": 2390,
	"createTime": 1530354842382,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 29,
		"code": "SYFX_ZC_RYCBZC",
		"name": "人员成本支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "7035.45",
	"indicatorIdAndBeginDate": 1529251200029
}, {
	"id": 2391,
	"createTime": 1530354842386,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 29,
		"code": "SYFX_ZC_RYCBZC",
		"name": "人员成本支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.202",
	"indicatorIdAndBeginDate": 1529251200029
}, {
	"id": 2392,
	"createTime": 1530354842390,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 29,
		"code": "SYFX_ZC_RYCBZC",
		"name": "人员成本支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200029
}, {
	"id": 2393,
	"createTime": 1530354842394,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 31,
		"code": "SYFX_ZC_NYFYZC",
		"name": "能源费用支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "1807.77",
	"indicatorIdAndBeginDate": 1529251200031
}, {
	"id": 2394,
	"createTime": 1530354842397,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 31,
		"code": "SYFX_ZC_NYFYZC",
		"name": "能源费用支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.052",
	"indicatorIdAndBeginDate": 1529251200031
}, {
	"id": 2395,
	"createTime": 1530354842401,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 31,
		"code": "SYFX_ZC_NYFYZC",
		"name": "能源费用支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200031
}, {
	"id": 2396,
	"createTime": 1530354842405,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 28,
		"code": "SYFX_ZC_FZZC",
		"name": "房租支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "16389.58",
	"indicatorIdAndBeginDate": 1529251200028
}, {
	"id": 2397,
	"createTime": 1530354842409,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 28,
		"code": "SYFX_ZC_FZZC",
		"name": "房租支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.470",
	"indicatorIdAndBeginDate": 1529251200028
}, {
	"id": 2398,
	"createTime": 1530354842413,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 28,
		"code": "SYFX_ZC_FZZC",
		"name": "房租支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200028
}, {
	"id": 2399,
	"createTime": 1530354842416,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 33,
		"code": "SYFX_ZC_FWFZC",
		"name": "服务费支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "3736.11",
	"indicatorIdAndBeginDate": 1529251200033
}, {
	"id": 2400,
	"createTime": 1530354842420,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 33,
		"code": "SYFX_ZC_FWFZC",
		"name": "服务费支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.107",
	"indicatorIdAndBeginDate": 1529251200033
}, {
	"id": 2401,
	"createTime": 1530354842424,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 33,
		"code": "SYFX_ZC_FWFZC",
		"name": "服务费支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200033
}, {
	"id": 2402,
	"createTime": 1530354842428,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 32,
		"code": "SYFX_ZC_QTZC",
		"name": "其他支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 22,
		"code": "SYFX_JE",
		"name": "金额"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "4770.27",
	"indicatorIdAndBeginDate": 1529251200032
}, {
	"id": 2403,
	"createTime": 1530354842432,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 32,
		"code": "SYFX_ZC_QTZC",
		"name": "其他支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 23,
		"code": "SYFX_BFZB",
		"name": "百分占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.137",
	"indicatorIdAndBeginDate": 1529251200032
}, {
	"id": 2404,
	"createTime": 1530354842435,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 32,
		"code": "SYFX_ZC_QTZC",
		"name": "其他支出",
		"level": 2,
		"parentId": 35
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200032
}, {
	"id": 2405,
	"createTime": 1530354842439,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 17,
		"code": "SYFX_YGZC",
		"name": "预估支出"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "34856.22",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2406,
	"createTime": 1530354842443,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 16,
		"code": "SYFX_SR",
		"name": "收入"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "40569.73",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2407,
	"createTime": 1530354842447,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 18,
		"code": "SYFX_YGLR",
		"name": "预估利润"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "5713.51",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2408,
	"createTime": 1530354842453,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 19,
		"code": "SYFX_LRMB",
		"name": "利润目标"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "8428.47",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2409,
	"createTime": 1530354842461,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 20,
		"code": "SYFX_QNLRMBDCL",
		"name": "全年利润目标达成率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.003",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2410,
	"createTime": 1530354842469,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 21,
		"code": "SYFX_QNTQ",
		"name": "去年同期"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "-",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2411,
	"createTime": 1530354842477,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 11,
		"code": "SYFX_ZRYGSY",
		"name": "昨日预估收益"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "5713.51",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2412,
	"createTime": 1530354842484,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 14,
		"code": "SYFX_BN",
		"name": "本年"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.00",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2413,
	"createTime": 1530354842488,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 15,
		"code": "SYFX_LJ",
		"name": "累计"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.00",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2414,
	"createTime": 1530354842492,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 13,
		"code": "SYFX_BYYG",
		"name": "本月预估"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "5713.51",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2415,
	"createTime": 1530354842495,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 36,
		"code": "SYFX_SYGK",
		"name": "收益概况",
		"level": 1,
		"parentId": 20
	},
	"demension": {
		"id": 12,
		"code": "SYFX_NHSYL",
		"name": "年化收益率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.002",
	"indicatorIdAndBeginDate": 1529251200036
}, {
	"id": 2416,
	"createTime": 1530354842500,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 300,
		"code": "RZLFX",
		"name": "入住率分析",
		"level": 0,
		"parentId": null
	},
	"demension": {
		"id": 100,
		"code": "RZLFX_JJ",
		"name": "均价"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "269",
	"indicatorIdAndBeginDate": 1529251200300
}, {
	"id": 2417,
	"createTime": 1530354842504,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 300,
		"code": "RZLFX",
		"name": "入住率分析",
		"level": 0,
		"parentId": null
	},
	"demension": {
		"id": 101,
		"code": "RZLFX_REVPAR",
		"name": "Revpar"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "271",
	"indicatorIdAndBeginDate": 1529251200300
}, {
	"id": 2418,
	"createTime": 1530354842508,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 300,
		"code": "RZLFX",
		"name": "入住率分析",
		"level": 0,
		"parentId": null
	},
	"demension": {
		"id": 102,
		"code": "RZLFX_REVPARZBB",
		"name": "Revpar指标比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "1.187",
	"indicatorIdAndBeginDate": 1529251200300
}, {
	"id": 2419,
	"createTime": 1530354842512,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 300,
		"code": "RZLFX",
		"name": "入住率分析",
		"level": 0,
		"parentId": null
	},
	"demension": {
		"id": 103,
		"code": "RZLFX_RZL",
		"name": "入住率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "1.006",
	"indicatorIdAndBeginDate": 1529251200300
}, {
	"id": 2420,
	"createTime": 1530354842516,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 300,
		"code": "RZLFX",
		"name": "入住率分析",
		"level": 0,
		"parentId": null
	},
	"demension": {
		"id": 104,
		"code": "RZLFX_RZLZBB",
		"name": "入住率指标比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "1.094",
	"indicatorIdAndBeginDate": 1529251200300
}, {
	"id": 2421,
	"createTime": 1530354842522,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 302,
		"code": "DDLYFX_ZYHY",
		"name": "中央会员",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 105,
		"code": "DDLYFX_JJ",
		"name": "均价"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "262",
	"indicatorIdAndBeginDate": 1529251200302
}, {
	"id": 2422,
	"createTime": 1530354842526,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 302,
		"code": "DDLYFX_ZYHY",
		"name": "中央会员",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 106,
		"code": "DDLYFX_JGZXL",
		"name": "价格执行率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.780",
	"indicatorIdAndBeginDate": 1529251200302
}, {
	"id": 2423,
	"createTime": 1530354842530,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 302,
		"code": "DDLYFX_ZYHY",
		"name": "中央会员",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 107,
		"code": "DDLYFX_RZLZB",
		"name": "入住率占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.172",
	"indicatorIdAndBeginDate": 1529251200302
}, {
	"id": 2424,
	"createTime": 1530354842534,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 303,
		"code": "DDLYFX_JBHY",
		"name": "家宾会员",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 105,
		"code": "DDLYFX_JJ",
		"name": "均价"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "294",
	"indicatorIdAndBeginDate": 1529251200303
}, {
	"id": 2425,
	"createTime": 1530354842538,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 303,
		"code": "DDLYFX_JBHY",
		"name": "家宾会员",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 106,
		"code": "DDLYFX_JGZXL",
		"name": "价格执行率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.876",
	"indicatorIdAndBeginDate": 1529251200303
}, {
	"id": 2426,
	"createTime": 1530354842542,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 303,
		"code": "DDLYFX_JBHY",
		"name": "家宾会员",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 107,
		"code": "DDLYFX_RZLZB",
		"name": "入住率占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.211",
	"indicatorIdAndBeginDate": 1529251200303
}, {
	"id": 2427,
	"createTime": 1530354842546,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 304,
		"code": "DDLYFX_XY",
		"name": "协议",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 105,
		"code": "DDLYFX_JJ",
		"name": "均价"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "286",
	"indicatorIdAndBeginDate": 1529251200304
}, {
	"id": 2428,
	"createTime": 1530354842550,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 304,
		"code": "DDLYFX_XY",
		"name": "协议",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 106,
		"code": "DDLYFX_JGZXL",
		"name": "价格执行率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.808",
	"indicatorIdAndBeginDate": 1529251200304
}, {
	"id": 2429,
	"createTime": 1530354842553,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 304,
		"code": "DDLYFX_XY",
		"name": "协议",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 107,
		"code": "DDLYFX_RZLZB",
		"name": "入住率占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.341",
	"indicatorIdAndBeginDate": 1529251200304
}, {
	"id": 2430,
	"createTime": 1530354842557,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 305,
		"code": "DDLYFX_ZJQD",
		"name": "中介渠道",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 105,
		"code": "DDLYFX_JJ",
		"name": "均价"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "269",
	"indicatorIdAndBeginDate": 1529251200305
}, {
	"id": 2431,
	"createTime": 1530354842561,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 305,
		"code": "DDLYFX_ZJQD",
		"name": "中介渠道",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 106,
		"code": "DDLYFX_JGZXL",
		"name": "价格执行率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.856",
	"indicatorIdAndBeginDate": 1529251200305
}, {
	"id": 2432,
	"createTime": 1530354842564,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 305,
		"code": "DDLYFX_ZJQD",
		"name": "中介渠道",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 107,
		"code": "DDLYFX_RZLZB",
		"name": "入住率占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.196",
	"indicatorIdAndBeginDate": 1529251200305
}, {
	"id": 2433,
	"createTime": 1530354842568,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 306,
		"code": "DDLYFX_SM",
		"name": "上门",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 105,
		"code": "DDLYFX_JJ",
		"name": "均价"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "275",
	"indicatorIdAndBeginDate": 1529251200306
}, {
	"id": 2434,
	"createTime": 1530354842571,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 306,
		"code": "DDLYFX_SM",
		"name": "上门",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 106,
		"code": "DDLYFX_JGZXL",
		"name": "价格执行率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.903",
	"indicatorIdAndBeginDate": 1529251200306
}, {
	"id": 2435,
	"createTime": 1530354842575,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 306,
		"code": "DDLYFX_SM",
		"name": "上门",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 107,
		"code": "DDLYFX_RZLZB",
		"name": "入住率占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.004",
	"indicatorIdAndBeginDate": 1529251200306
}, {
	"id": 2436,
	"createTime": 1530354842578,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 307,
		"code": "DDLYFX_SWRF",
		"name": "商务日房",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 105,
		"code": "DDLYFX_JJ",
		"name": "均价"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "87",
	"indicatorIdAndBeginDate": 1529251200307
}, {
	"id": 2437,
	"createTime": 1530354842581,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 307,
		"code": "DDLYFX_SWRF",
		"name": "商务日房",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 106,
		"code": "DDLYFX_JGZXL",
		"name": "价格执行率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.290",
	"indicatorIdAndBeginDate": 1529251200307
}, {
	"id": 2438,
	"createTime": 1530354842585,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 307,
		"code": "DDLYFX_SWRF",
		"name": "商务日房",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 107,
		"code": "DDLYFX_RZLZB",
		"name": "入住率占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.058",
	"indicatorIdAndBeginDate": 1529251200307
}, {
	"id": 2439,
	"createTime": 1530354842588,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 308,
		"code": "DDLYFX_QT",
		"name": "其他",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 105,
		"code": "DDLYFX_JJ",
		"name": "均价"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "272",
	"indicatorIdAndBeginDate": 1529251200308
}, {
	"id": 2440,
	"createTime": 1530354842591,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 308,
		"code": "DDLYFX_QT",
		"name": "其他",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 106,
		"code": "DDLYFX_JGZXL",
		"name": "价格执行率"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.670",
	"indicatorIdAndBeginDate": 1529251200308
}, {
	"id": 2441,
	"createTime": 1530354842595,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 308,
		"code": "DDLYFX_QT",
		"name": "其他",
		"level": 1,
		"parentId": 301
	},
	"demension": {
		"id": 107,
		"code": "DDLYFX_RZLZB",
		"name": "入住率占比"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0.023",
	"indicatorIdAndBeginDate": 1529251200308
}, {
	"id": 2442,
	"createTime": 1530354842600,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 309,
		"code": "BKSFX",
		"name": "办卡数分析",
		"level": 0,
		"parentId": null
	},
	"demension": {
		"id": 108,
		"code": "BKSFX_JKS",
		"name": "金卡数"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0",
	"indicatorIdAndBeginDate": 1529251200309
}, {
	"id": 2443,
	"createTime": 1530354842603,
	"merchant": {
		"id": 1,
		"name": "上海源意酒店有限公司"
	},
	"indicator": {
		"id": 309,
		"code": "BKSFX",
		"name": "办卡数分析",
		"level": 0,
		"parentId": null
	},
	"demension": {
		"id": 109,
		"code": "BKSFX_PKS",
		"name": "普卡数"
	},
	"dataPacket": {
		"id": 90,
		"filePath": null,
		"up2Chain": false
	},
	"beginDate": 1529251200000,
	"endDate": 1529337600000,
	"period": "d",
	"value": "0",
	"indicatorIdAndBeginDate": 1529251200309
}]`
