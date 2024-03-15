package xrobot

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

/**



Created on : 2018/2/19

*/

var AES = DefaultAES{}

type DefaultAES struct {
}

const DefaultAESKey = "12345678910Asdfx"
const IsOpenAES = true

// 加密
func (x DefaultAES) AesEncryptBean(bean interface{}) string {
	jsonBytes, err := json.Marshal(bean)
	if err != nil {
		panic(bean)
	}
	if !IsOpenAES {
		return string(jsonBytes)
	}
	dataBytes, err := x.AesEncrypt([]byte(jsonBytes), []byte(DefaultAESKey))
	retStr := ""
	if dataBytes != nil {
		retStr = base64.StdEncoding.EncodeToString(dataBytes)
	}
	if err != nil {
		retStr = err.Error()
	}
	return retStr
}

// 加密
func (x DefaultAES) AesEncryptStr(origData string) string {
	if !IsOpenAES {
		return origData
	}
	dataBytes, err := x.AesEncrypt([]byte(origData), []byte(DefaultAESKey))
	retStr := ""
	if dataBytes != nil {
		retStr = base64.StdEncoding.EncodeToString(dataBytes)
	}
	if err != nil {
		retStr = err.Error()
	}
	return retStr
}

func (x DefaultAES) AesEncryptStr_Must(origData string) string {
	dataBytes, err := x.AesEncrypt([]byte(origData), []byte(DefaultAESKey))
	retStr := ""
	if dataBytes != nil {
		retStr = base64.StdEncoding.EncodeToString(dataBytes)
	}
	if err != nil {
		retStr = err.Error()
	}
	return retStr
}

func (x DefaultAES) AesDecryptStr_Must(encrypted string) string {
	// 验证输入参数
	// 必须为aes.BlockSize的倍数 --- 16
	encBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if len(encBytes)%aes.BlockSize != 0 || err != nil {
		return "invalid"
	}
	dataBytes, err := x.AesDecrypt(encBytes, []byte(DefaultAESKey))
	retStr := ""
	if dataBytes != nil {
		retStr = string(dataBytes)
	}
	if err != nil {
		retStr = err.Error()
	}
	return retStr
}

// 解密
func (x DefaultAES) AesDecryptStr(encrypted string) string {
	// 验证输入参数
	// 必须为aes.BlockSize的倍数 --- 16
	if !IsOpenAES {
		return encrypted
	}
	encBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if len(encBytes)%aes.BlockSize != 0 || err != nil {
		return "解密错误"
	}
	dataBytes, err := x.AesDecrypt(encBytes, []byte(DefaultAESKey))
	retStr := ""
	if dataBytes != nil {
		retStr = string(dataBytes)
	}
	if err != nil {
		retStr = err.Error()
	}
	return retStr
}

func (x DefaultAES) AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (x DefaultAES) AesDecrypt(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(encrypted))
	// origData := encrypted
	blockMode.CryptBlocks(origData, encrypted)
	origData, err = pKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	if err != nil {
		return origData, err
	}
	return origData, nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize //需要padding的数目
	//只要少于256就能放到一个byte中，默认的blockSize=16(即采用16*8=128, AES-128长的密钥)
	//最少填充1个byte，如果原文刚好是blocksize的整数倍，则再填充一个blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding) //生成填充的文本
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length < 1 {
		return []byte{}, errors.New("加/解密错误")
	}
	unpadding := int(origData[length-1])
	tarLen := length - unpadding
	if length < tarLen || tarLen < 0 {
		return []byte{}, errors.New("加/解密错误")
	}
	//fmt.Println(length,unpadding,tarLen)
	return origData[:tarLen], nil
}
