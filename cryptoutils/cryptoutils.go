package cryptoutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/gob"
	"fmt"

	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/thomasks/eju-chaincodes/kmc"
)

//CryptoDescriptor comments
type CryptoDescriptor struct {
	Level        string   `json:"level"`
	CryptoFields []string `json:"cryptoFields"`
}

//CryptoDataByDescriptor export
func CryptoDataByDescriptor(stub shim.ChaincodeStubInterface, rawData map[string]interface{}, cds []CryptoDescriptor) {
	for _, cd := range cds {
		cryptokey := kmc.GetCryptoKey(stub, cd.Level)
		keys := cd.CryptoFields
		for _, key := range keys {
			rawValue := rawData[key]
			rawBytes, err := getBytes(rawValue)
			if err == nil {
				encryptValue, err2 := encryptData(rawBytes, cryptokey)
				if err2 == nil {
					rawData[key] = encryptValue
				}
			}
		}
	}
}

//DecryptoDataByDescriptor export
func DecryptoDataByDescriptor(stub shim.ChaincodeStubInterface, encryptData map[string]interface{}, cds []CryptoDescriptor) {
	for _, cd := range cds {
		cryptokey := kmc.GetCryptoKey(stub, cd.Level)
		keys := cd.CryptoFields
		for _, key := range keys {
			encryptValue, _ := encryptData[key].([]byte)
			rawValueBytes, err := decryptData(encryptValue, cryptokey)

			if err == nil {
				var obj = new(interface{})
				err2 := getObj(rawValueBytes, obj)
				if err2 == nil {
					encryptData[key] = obj
				}
			}
		}
	}
}

func encryptData(data []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	b := data
	ciphertext := make([]byte, aes.BlockSize+len(b))
	fmt.Printf("blockSize=%d,cipherText=%s\n", aes.BlockSize, string(ciphertext))
	iv := ciphertext[:aes.BlockSize]
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return []byte(base64.StdEncoding.EncodeToString(ciphertext)), nil
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getObj(dataBytes []byte, data *interface{}) error {
	var buf *bytes.Buffer
	buf = bytes.NewBuffer(dataBytes)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(data)
	return err
}

func decryptData(text []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
