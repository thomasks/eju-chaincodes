package kmc

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//
var (
	cryptoDict = map[string]string{
		"PUBLIC":  "01FtVvUe87kbSrBwzj2fx8tHGcdREIUH",
		"ORG":     "wRG7GDJMPavbHJjDsNNU0Tlv7ivO8Nic",
		"GROUP":   "LHmYrRJ8UCDDks4qQ3XhYj4Ty9HVGWqi",
		"PRIVATE": "leoi9HG4XJEz68Hc5r9KK1N4zCdTP1wI",
	}
)

//GetCryptoKey exported
func GetCryptoKey(stub shim.ChaincodeStubInterface, level string) string {
	/*switch level {
	case "PUBLIC":
		cryptoKey = ""
	case "ORG":
		cryptoKey = ""
	case "GROUP":
		cryptoKey = ""
	case "PRIVATE":
		cryptoKey = ""
	}*/
	return cryptoDict[level]
}
